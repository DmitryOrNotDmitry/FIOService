package logger

import (
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func Init(level logrus.Level) {
	Log = logrus.New()
	Log.SetLevel(level)
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
