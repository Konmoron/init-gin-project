package controllers

import (
	"net/http"
	"time"

	"github.com/Konmoron/init-gin-project/config"
	"github.com/gin-gonic/gin"
)

// Healthcheck : Healthcheck
func Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"app_name": config.Cfg.AppName,
		"code":     200,
		"status":   "ok",
		"time":     time.Now(),
	})
}
