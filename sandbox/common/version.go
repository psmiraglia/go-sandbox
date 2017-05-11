package common

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	major = 0
	minor = 1
	patch = 0
	release = ""
	build = ""
	commit = ""
)

func Version() string {
	if len(release) > 0 {
		return fmt.Sprintf(
			"%d.%d.%d_%s (build:%s, commit:%s)",
			major, minor, patch, release, build, commit,
		)
	} else {
		return fmt.Sprintf(
			"%d.%d.%d (build:%s, commit:%s)",
			major, minor, patch, build, commit,
		)
	}
}

func NamedVersion() string {
	return fmt.Sprintf(
		"%s-%s",
		filepath.Base(os.Args[0]), Version(),
	)
}
