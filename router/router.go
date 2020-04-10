package router

import (
	"time"

	"github.com/Konmoron/init-gin-project/config"
	"github.com/Konmoron/init-gin-project/controllers"
	"github.com/Konmoron/init-gin-project/log"

	"github.com/gin-contrib/pprof"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

// InitRouter : init router
func InitRouter() *gin.Engine {
	// set gin model
	gin.SetMode(config.Cfg.GinModel)

	router := gin.New()

	// router.Use(ginzap.Ginzap(log.GinLogger, time.RFC3339, true))
	// router.Use(ginzap.Ginzap(log.GinLogger, "2006-01-02 15:04:05.000", false))
	router.Use(ginzap.Ginzap(log.GinLogger, time.RFC3339, false))
	router.Use(ginzap.RecoveryWithZap(log.GinLogger, true))

	// set healthcheck and / location
	router.GET("/healthcheck", controllers.Healthcheck)
	router.GET("/", controllers.Healthcheck)

	v1 := router.Group("/v1")
	{
		v1.GET("/healthcheck", controllers.Healthcheck)
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/healthcheck", controllers.Healthcheck)
	}

	// set pprof
	pprof.Register(router, "/debug/pprof")

	return router
}
