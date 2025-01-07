package version_test

import (
	"fmt"
	"runtime"
	"testing"

	usercase "github.com/danielsoro/wordpress-cli/lib/version"
)

func TestVersionCommand(t *testing.T) {
	version := usercase.NewVersion()
	versionOutput := version.Execute()
	expect := fmt.Sprintf("wordpress-cli version 0.0.1-SNAPSHOT %s/%s", runtime.GOOS, runtime.GOARCH)

	if versionOutput != expect {
		t.Fatalf("Expecting %s, but get %s", expect, versionOutput)
	}
}
