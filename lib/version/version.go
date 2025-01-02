package version

import (
	"fmt"
	"runtime"
)

const VERSION = "0.0.1-SNAPSHOT"

type Version struct {
	version string
}

func NewVersion() Version {
	return Version{
		version: VERSION,
	}
}

func (v *Version) Execute() string {
	return fmt.Sprintf("wordpress-cli version %s %s/%s", VERSION, runtime.GOOS, runtime.GOARCH)
}
