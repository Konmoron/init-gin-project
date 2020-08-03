package main

import (
	"github.com/Konmoron/init-gin-project/config"
	"github.com/Konmoron/init-gin-project/log"
	"github.com/Konmoron/init-gin-project/router"
)

func init() {
	config.InitConf("config/config.yaml")

	log.InitLog(config.Cfg.LogLevel, "./log")

	log.MainLogger.Info("successfully init!")
}

// @title init-gin-project
// @version 1.0
// @description init-gin-project

// @contact.name Konmoron
// @contact.email cqlq2012@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080

func main() {
	r := router.InitRouter()

	r.Run(config.Cfg.Addr)
}
