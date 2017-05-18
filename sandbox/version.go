package sandbox

import (
	"fmt"
)

var (
	major = 0
	minor = 1
	patch = 0
	release = ""
)

var (
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
