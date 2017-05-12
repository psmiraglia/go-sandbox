package logging

import (
	"github.com/sirupsen/logrus"
	"github.com/psmiraglia/go-sandbox/sandbox/version"
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
	"@build": version.Build,
	"@commit": version.Commit,
	"@version": 1,
})
