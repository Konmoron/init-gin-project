package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// MainLogger : main logger
var MainLogger *zap.Logger

// GinLogger : gin log
var GinLogger *zap.Logger

// InitLog : init all logger
func InitLog(levelString string) {
	// set log level
	var level zapcore.Level
	err := level.Set(levelString)
	if err != nil {
		level.Set("INFO")
	}

	// set file path
	MainLogger = NewLogger(
		"./log/main.log",
		level,
		1024,
		30,
		1,
		true,
		"main")
	// set file path
	GinLogger = NewLogger(
		"./log/gin.log",
		level,
		1024,
		30,
		1,
		true,
		"gin")
	MainLogger.Info("start logs successfully")
}
