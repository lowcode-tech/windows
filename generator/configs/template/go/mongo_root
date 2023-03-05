package app

import (
	"github.com/core-go/core"
	mid "github.com/core-go/log/middleware"
	"github.com/core-go/log/zap"
	"github.com/core-go/mongo"
)

type Config struct {
	Server     core.ServerConf    `mapstructure:"server"`
	Mongo      mongo.MongoConfig  `mapstructure:"mongo"`
	Log        log.Config         `mapstructure:"log"`
	MiddleWare mid.LogConfig      `mapstructure:"middleware"`
	Status     *core.StatusConfig `mapstructure:"status"`
	Action     *core.ActionConfig `mapstructure:"action"`
}
