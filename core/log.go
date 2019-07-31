package core

import (
	"flag"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLog() {
	Log = logrus.New()
	level := flag.String("log-level", "info", "log level")
	flag.Parse()
	switch *level {
	case "trace":
		Log.SetLevel(logrus.TraceLevel)
	case "debug":
		Log.SetLevel(logrus.DebugLevel)
	case "info":
		Log.SetLevel(logrus.InfoLevel)
	case "warning":
		Log.SetLevel(logrus.WarnLevel)
	case "warn":
		Log.SetLevel(logrus.WarnLevel)
	case "error":
		Log.SetLevel(logrus.ErrorLevel)
	case "fatal":
		Log.SetLevel(logrus.FatalLevel)
	case "panic":
		Log.SetLevel(logrus.PanicLevel)
	default:
		Log.SetLevel(logrus.InfoLevel)
	}
}
