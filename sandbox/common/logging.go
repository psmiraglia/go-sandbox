package common

import (
	"github.com/sirupsen/logrus"
	"time"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		FieldMap: logrus.FieldMap {
			logrus.FieldKeyTime: "@timestamp",
		},
	})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.AddHook(NewMyHook())
}

var Log = logrus.WithFields(logrus.Fields{
	"@build": Build,
	"@commit": Commit,
	"@version": 1,
})
