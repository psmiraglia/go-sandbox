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

// These variables are set at build time
var (
	// The date at build time
	build = ""

	// The git commit at build time
	commit = ""
)

// Version returns the project's version as string.
func Version() string {
	if len(release) > 0 {
		return fmt.Sprintf(
			"%d.%d.%d_%s (build:%s, commit:%s)",
			major, minor, patch, release, build, commit,
		)
	}
	return fmt.Sprintf(
		"%d.%d.%d (build:%s, commit:%s)",
		major, minor, patch, build, commit,
	)
}
