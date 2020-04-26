package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Konmoron/init-gin-project/config"
	"github.com/Konmoron/init-gin-project/log"
	"github.com/Konmoron/init-gin-project/router"
	"github.com/Konmoron/init-gin-project/utils"
	"github.com/stretchr/testify/assert"
)

// init unittest env
func init() {
	// init config
	config.InitConf("../config/config.yaml")

	// init log
	log.InitLog(config.Cfg.LogLevel, "../log")
}

func TestHealthcheckRoute(t *testing.T) {
	r := router.InitRouter()

	res, body := utils.UnittestGet("/healthcheck", r)

	var resData map[string]interface{}

	if err := json.Unmarshal(body, &resData); err != nil {
		fmt.Println("json.unmarshal(body, &response) err: ", err)
	}

	// print resData
	// t.Log(resData)

	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, "ok", resData["status"], "status should be ok")
	// this will be error
	// assert.Equal(t, "ok1", response["status"], "status should be ok")
}
