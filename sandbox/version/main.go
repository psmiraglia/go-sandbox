package version

import (
	"fmt"
	"github.com/psmiraglia/go-sandbox/sandbox/common"
)

var (
	major = 0
	minor = 1
	patch = 0
	release = ""
)

var (
	Build = ""
	Commit = ""
)

func Version() string {
	if len(release) > 0 {
		return fmt.Sprintf(
			"%d.%d.%d_%s (build:%s, commit:%s)",
			major, minor, patch, release, Build, Commit,
		)
	} else {
		return fmt.Sprintf(
			"%d.%d.%d (build:%s, commit:%s)",
			major, minor, patch, Build, Commit,
		)
	}
}

func NamedVersion() string {
	return fmt.Sprintf(
		"%s-%s",
		common.Name, Version(),
	)
}
