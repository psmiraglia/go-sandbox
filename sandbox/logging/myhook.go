package logging

import (
	"crypto/sha512"
	"github.com/psmiraglia/go-sandbox/sandbox/common"
	"github.com/sirupsen/logrus"
	"io"
)

type myHook struct {
    // None
}

func NewMyHook() *myHook {
	return &myHook{}
}

func (hook *myHook) Fire(e *logrus.Entry) error {
	hash := sha512.New()
	io.WriteString(hash, e.Message)
	io.WriteString(hash, e.Time.Format(common.TimestampFormat))
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
