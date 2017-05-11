package common

import (
	"fmt"
)

var Name = "sandbox"
var major = 0
var minor = 1
var patch = 0
var release = ""

func Version() string {
	if len(release) > 0 {
		return fmt.Sprintf("%d.%d.%d_%s", major, minor, patch, release)
	} else {
		return fmt.Sprintf("%d.%d.%d", major, minor, patch)
	}
}

func NamedVersion() string {
	return fmt.Sprintf("%s-%s", Name, Version())
}
