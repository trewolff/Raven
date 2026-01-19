package logging

import (
	"os"
	"raven/internal/config"

	"github.com/sirupsen/logrus"
	logrusadapter "logur.dev/adapter/logrus"
	"logur.dev/logur"
)

var Logger logur.KVLoggerFacade

func InitLogger() {
	conf, _ := config.GetConfig()
	logrusAdapter := logrus.New()
	logrusAdapter.SetOutput(os.Stdout)
	logrusAdapter.Formatter = &logrus.JSONFormatter{}
	logrusAdapter.Level = conf.LOG_LEVEL
	Logger = logur.LoggerToKV(logrusadapter.New(logrusAdapter))
}
