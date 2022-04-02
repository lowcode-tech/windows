package app

import (
	"github.com/core-go/log"
	mid "github.com/core-go/log/middleware"
	"github.com/core-go/mongo"
	sv "github.com/core-go/service"
)

type Config struct {
	Server     sv.ServerConf     `mapstructure:"server"`
	Mongo      mongo.MongoConfig `mapstructure:"mongo"`
	Log        log.Config        `mapstructure:"log"`
	MiddleWare mid.LogConfig     `mapstructure:"middleware"`
	Status     *sv.StatusConfig  `mapstructure:"status"`
	Action     *sv.ActionConfig  `mapstructure:"action"`
}
