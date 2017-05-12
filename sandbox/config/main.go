package config

import (
	"github.com/psmiraglia/go-sandbox/sandbox/common"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"fmt"
)

var config = viper.New()

func init() {
	config.SetConfigName("sandbox")
	config.AddConfigPath(".")
	err := config.ReadInConfig()
	if err != nil {
		fmt.Println("Fatal error config file: %s", err)
	}

	config.SetDefault("log.level", "info")
	config.SetDefault("log.format", "text")
}

func ParseLogLevel() logrus.Level {
    switch level := config.GetString("log.level"); level {
        case "debug":
            return logrus.DebugLevel
        case "info":
            return logrus.InfoLevel
		case "warn":
			return logrus.WarnLevel
		case "error":
			return logrus.ErrorLevel
		case "fatal":
			return logrus.FatalLevel
        default:
            return logrus.InfoLevel
    }
}

func ParseLogFormat() logrus.Formatter {
    switch format := config.GetString("log.format"); format {
        case "json":
            return &logrus.JSONFormatter {
				TimestampFormat: common.TimestampFormat,
				FieldMap: logrus.FieldMap {
					logrus.FieldKeyTime: "@timestamp",
					logrus.FieldKeyMsg:  "message",
				},
            }
		case "text":
			return &logrus.TextFormatter {
				TimestampFormat: common.TimestampFormat,
                FullTimestamp: true,
            }
        default:
            return &logrus.TextFormatter {
				TimestampFormat: common.TimestampFormat,
                FullTimestamp: true,
            }
    }
}
