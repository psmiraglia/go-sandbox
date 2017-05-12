package logging

import (
	"github.com/psmiraglia/go-sandbox/sandbox/config"
	"github.com/psmiraglia/go-sandbox/sandbox/version"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	// set output format from config file
    logrus.SetFormatter(config.ParseLogFormat())
	// set level from config file
	logrus.SetLevel(config.ParseLogLevel())
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
	"build":  version.Build,
	"commit": version.Commit,
})
