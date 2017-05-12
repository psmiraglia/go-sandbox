package logging

import (
	"github.com/psmiraglia/go-sandbox/sandbox/common"
	"github.com/psmiraglia/go-sandbox/sandbox/version"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: common.TimestampFormat,
		FieldMap: logrus.FieldMap {
			logrus.FieldKeyTime: "@timestamp",
            logrus.FieldKeyMsg:  "message",
		},
	})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	logrus.AddHook(NewMyHook())
}

// Logger object to be included in all the source files to produce log
// entries.
//
//  var log = logging.Log
//  ...
//  log.Info("This is a info entry")
var Log = logrus.WithFields(logrus.Fields{
	"build":    version.Build,
	"commit":   version.Commit,
	"@version": 1,
})
