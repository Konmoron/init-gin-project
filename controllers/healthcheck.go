package controllers

import (
	"net/http"
	"time"

	"github.com/Konmoron/init-gin-project/config"
	"github.com/Konmoron/init-gin-project/model"
	"github.com/gin-gonic/gin"
)

// Healthcheck : Healthcheck
// @Summary healthcheck
// @description healthcheck
// @description
// @version 1.0
// @Success 200 {object} model.HealthcheckResult healthcheck result
// @Router /healthcheck [get]
func Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, model.HealthcheckResult{
		AppName: config.Cfg.AppName,
		Code:    http.StatusOK,
		Status:  "ok",
		Time:    time.Now(),
	})
}
