package sandbox

import (
	"github.com/sirupsen/logrus"
	"os"
)

const (
	// TimestampFormat is the timestamp format to be used
    TimestampFormat = "2006-01-02T15:04:05.000Z07:00"
)

// Logger is the default logger
var Logger = logrus.New()

func init() {
	cfg := NewConfig()

	formatter := cfg.LogFormat
	level := cfg.LogLevel

	Logger.Out = os.Stdout
	Logger.Formatter = formatter
	Logger.Level = level
}
