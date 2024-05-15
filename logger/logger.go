package logger

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	level, err := logrus.ParseLevel("error")
	if err != nil {
		log.Fatal("logrus.ParseLevel error! ", err)
	}
	Logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02T15:04:05.000Z07:00"})
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(level)
}
