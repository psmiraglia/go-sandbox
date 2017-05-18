package sandbox

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Configuration directives
const (
	logFormat = "log.format"
	logLevel = "log.level"
)

// Config is a configuration registry. It contains an instance of the
// configuration management engine (Viper) and a set of configuration
// parameters loaded from a configuration file.
type Config struct {
	// Viper instance
	v         *viper.Viper

	// LogLevel to be considered as minimum log severity level
	LogLevel  logrus.Level

	// LogFormat to be used when log entries are generated
	LogFormat logrus.Formatter
}

// NewConfig returns an initialised Config instance.
func NewConfig() *Config {
	log := Logger

	cfg := new(Config)
	cfg.v = viper.New()
	cfg.v.SetConfigName("sandbox")
	cfg.v.AddConfigPath(".")

	err := cfg.v.ReadInConfig()
	if err != nil {
		log.Warn(err)
		log.Info("Using default configuration values")
	}

	cfg.v.SetDefault(logFormat, "text")
	cfg.v.SetDefault(logLevel, "info")

	cfg.parseLogFormat()
    cfg.parseLogLevel()

	return cfg
}

// parseLogFormat parses log.format configuration directive.
func (cfg *Config) parseLogFormat() {
	log := Logger
	switch format := cfg.v.GetString(logFormat); format {
		case "json":
			cfg.LogFormat = &logrus.JSONFormatter {
				TimestampFormat: TimestampFormat,
				FieldMap: logrus.FieldMap {
					logrus.FieldKeyMsg:  "message",
				},
			}
		case "text":
			cfg.LogFormat = &MyTextFormatter {}
		default:
			log.Warn("Parsed value not managed for ", logFormat)
			cfg.LogFormat = &logrus.TextFormatter {
				TimestampFormat: TimestampFormat,
				FullTimestamp: true,
			}

	}
}

// parseLogFormat parses log.level configuration directive.
func (cfg *Config) parseLogLevel() {
	log := Logger
	switch level := cfg.v.GetString(logLevel); level {
		case "debug":
			cfg.LogLevel = logrus.DebugLevel
		case "info":
			cfg.LogLevel = logrus.InfoLevel
		case "warn":
			cfg.LogLevel = logrus.WarnLevel
		case "error":
			cfg.LogLevel = logrus.ErrorLevel
		case "fatal":
			cfg.LogLevel = logrus.FatalLevel
		default:
			log.Warn("Parsed value not managed for ", logLevel)
			cfg.LogLevel = logrus.InfoLevel
	}
}
