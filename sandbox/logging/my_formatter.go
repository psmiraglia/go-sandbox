package logging

import (
	"github.com/psmiraglia/go-sandbox/sandbox/common"
	"github.com/sirupsen/logrus"
	"bytes"
)

type MyTextFormatter struct {
	// Prefix string according to the log severity.
	Prefix string
}

// SetPrefix sets the log entry prefix according to the entry's severity.
func (f *MyTextFormatter) setPrefix(e *logrus.Entry) {
	switch level := e.Level.String(); level {
		case "debug":
			f.Prefix = "[.] "
		case "info":
			f.Prefix = "[+] "
		case "warn":
			f.Prefix = "[!] "
		case "error":
			f.Prefix = "[x] "
		case "fatal":
			f.Prefix = "[P] "
		default:
			f.Prefix = "[*] "
	}
}

// Format formats log entries according to the following rule
//
//		<prefix> <timestamp> {MYFMT} <message>
//
// <prefix> follows the log severity and it is set by setPrefix() function
// while {MYFMT} is a fixed string.
func (f *MyTextFormatter) Format(e *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if e.Buffer != nil {
		b = e.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	f.setPrefix(e)
	b.WriteString(f.Prefix)
	b.WriteString(e.Time.Format(common.TimestampFormat))
	b.WriteString(" {MYFMT} ")
	b.WriteString(e.Message)
	b.WriteByte('\n')
	return b.Bytes(), nil
}
