package model

import "time"

// HealthcheckResult healthcheck result
type HealthcheckResult struct {
	AppName string    `json:"app_name"`
	Code    int       `json:"code"`
	Status  string    `json:"status"`
	Time    time.Time `json:"time"`
}
