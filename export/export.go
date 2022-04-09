package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"os"
	"path/filepath"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-generator/core"
	"github.com/go-generator/core/build"
	"github.com/go-generator/core/export"
	gdb "github.com/go-generator/core/export/db"
	"github.com/go-generator/core/export/relationship"
	"github.com/go-generator/core/export/types"
	"github.com/go-generator/core/generator"
	"github.com/go-generator/core/io"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	var prj *metadata.Project
	name := "models.json"
	prName := "project.json"
	ctx := context.TODO()
	arguments := os.Args
	if len(arguments) < 1 {
		log.Println("error: not enough arguments")
		return
	}
	dsn := arguments[1]
	driver := gdb.DetectDriver(dsn)
	dbName, err := gdb.ExtractDBName(dsn, driver)
	if err != nil {
		log.Println("error: ", err)
		return
	}
	sqlDb, err := sql.Open(driver, dsn)
	if err != nil {
		log.Println("error: ", err)
		return
	}
	defer func(sDB *sql.DB) {
		err := sDB.Close()
		if err != nil {
			log.Println("error: ", err)
		}
	}(sqlDb)

	tables, err := gdb.ListTables(ctx, sqlDb, dbName)
	if err != nil {
		log.Println("error: ", err)
		return
	}

	primaryKeys, err := export.GetAllPrimaryKeys(ctx, sqlDb, dbName, driver, tables)
	if err != nil {
		log.Println("error: ", err)
		return
	}

	rl, err := relationship.GetRelationshipTable(ctx, sqlDb, dbName, tables, primaryKeys)
	if err != nil {
		log.Println("error: ", err)
		return
	}

	t, ok := types.Types[driver]
	if !ok {
		log.Println("error: ", "universal type conversion not found")
		return
	}
	models, err := export.ToModels(ctx, sqlDb, dbName, tables, rl, t, primaryKeys)
	if err != nil {
		log.Println("error: ", err)
		return
	}

	option := flag.NewFlagSet("", flag.ExitOnError)
	f := option.String("f", "", "output name")
	p := option.String("p", "go_sql", "project type")
	m := option.String("m", "go-service", "go module")
	prjTmp := option.String("pt", filepath.Join(".", "project"), "project template")
	if len(arguments) > 2 {
		err := option.Parse(arguments[2:])
		if err != nil {
			log.Println("error: ", err)
			return
		}
	}

	if *f != "" {
		name = *f
	}
	if *p == "" {
		err = io.SaveModels(filepath.Join(".", name), models, true)
		if err != nil {
			log.Println("error: ", err)
			return
		}
	} else {
		prjPath, err := filepath.Abs(filepath.Join(".", "project"))
		if err != nil {
			log.Println("error: ", err)
			return
		}
		exist, err := exists(prjPath)
		if err != nil {
			log.Fatalln(err)
		}
		if !exist {
			prjPath, err = filepath.Abs(filepath.Join(".", "configs", "project"))
			if err != nil {
				log.Fatalln(err)
			}
			exist, err = exists(prjPath)
			if err != nil {
				log.Fatalln(err)
			}
			if !exist {
				prjPath, err = filepath.Abs(*prjTmp)
				if err != nil {
					log.Fatalln(err)
				}
				exist, err = exists(prjPath)
				if err != nil {
					log.Fatalln(err)
				}
				if !exist {
					log.Fatal("invalid project template path")
				}
			}
		}

		projectTemplate, err := io.Load(prjPath)
		if err != nil {
			log.Fatalln(err)
		}
		prj, err = generator.ExportProject(*p, name, projectTemplate, models, build.InitEnv)
		if err != nil {
			err = io.SaveModels(filepath.Join(".", name), models, true)
			if err != nil {
				log.Println("error: ", err)
				return
			}
			return
		}
		_, ok := prj.Env["go_module"]
		if ok {
			prj.Env["go_module"] = *m
		}
		if *f != "" {
			prName = *f
		}
		err = io.SaveProject(filepath.Join(".", prName), *prj, true)
		if err != nil {
			log.Println("error: ", err)
			return
		}
	}
}
