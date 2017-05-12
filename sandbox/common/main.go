// Package common defines common variables to be used in the project.
package common

import (
	"path/filepath"
    "os"
)

// It stores the application name.
var Name = filepath.Base(os.Args[0])

// It stores the timestamp format to be used in log entries.
var TimestampFormat = "2006-01-02T15:04:05.000Z07:00"
