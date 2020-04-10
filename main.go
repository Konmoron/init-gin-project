package main

import (
	"github.com/Konmoron/init-gin-project/config"
	"github.com/Konmoron/init-gin-project/log"
	"github.com/Konmoron/init-gin-project/router"
)

func init() {
	config.InitConf("config/config.yaml")

	log.InitLog(config.Cfg.LogLevel)

	log.MainLogger.Info("successfully init!")
}

func main() {
	r := router.InitRouter()

	r.Run(config.Cfg.Addr)
}
