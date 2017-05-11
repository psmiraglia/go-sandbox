package common

import (
	//"github.com/psmiraglia/go-sandbox/sandbox/common"
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
}

var Log = logrus.WithFields(logrus.Fields{
	"@build": Build,
	"@commit": Commit,
	"@version": 1,
})
