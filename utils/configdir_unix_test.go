package configdir_test

import (
	"testing"

	configdir "github.com/danielsoro/wordpress-cli/utils"
)

func TestLocalConfig(t *testing.T) {
	expect := "/home/dcunha/.config/wordpress-cli/config"

	config := configdir.LocalConfig("wordpress-cli/config")
	if config != expect {
		t.Fatalf("Expected %s, got %s", expect, config)
	}
}
