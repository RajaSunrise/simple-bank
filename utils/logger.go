package utils

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()

	Log.SetOutput(&lumberjack.Logger{
		Filename: "app/log.app",
		MaxSize: 10,
		MaxBackups: 3,
		MaxAge: 30,
	})

	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors: true,
		ForceQuote: true,
	})
}
