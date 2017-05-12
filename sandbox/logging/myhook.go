package logging

import (
	"crypto/sha256"
	"github.com/sirupsen/logrus"
	"io"
	"time"
)

type myHook struct {
    // None
}

func NewMyHook() *myHook {
	return &myHook{}
}

func (hook *myHook) Fire(e *logrus.Entry) error {
	hash := sha256.New()
	io.WriteString(hash, e.Message)
	io.WriteString(hash, e.Time.Format(time.RFC3339Nano))
	e.Data["hash"] = hash.Sum(nil)
	return nil
}

func (hook *myHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}
