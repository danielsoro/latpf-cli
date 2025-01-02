package version_test

import (
	"testing"

	usercase "github.com/danielsoro/wordpress-cli/lib/version"
)

func TestVersionCommand(t *testing.T) {
	version := usercase.NewVersion()
	versionOutput := version.Execute()
	expect := "wordpress-cli version 0.0.1-SNAPSHOT linux/amd64"

	if versionOutput != expect {
		t.Fatalf("Expecting %s, but get %s", expect, versionOutput)
	}
}
