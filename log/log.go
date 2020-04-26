package log

import (
	"path"

	"go.uber.org/zap"
)

// MainLogger : main logger
var MainLogger *zap.Logger

// GinLogger : gin log
var GinLogger *zap.Logger

// InitLog : init all logger
// levelString log level string: "DEBUG" "INFO" ...
// logDir log file dir
func InitLog(levelString, logDir string) {
	MainLogger = NewZapLoggerSimple(levelString, path.Join(logDir, "main.log"), "main")
	GinLogger = NewZapLoggerSimple(levelString, path.Join(logDir, "gin.log"), "gin")
	MainLogger.Info("start logs successfully")
}
