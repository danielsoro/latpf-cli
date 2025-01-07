package configdir_test

import (
	"fmt"
	"os"
	"testing"

	configdir "github.com/danielsoro/wordpress-cli/utils"
)

func TestLocalConfig(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}

	expect := fmt.Sprintf("%s/.config/wordpress-cli/config", home)

	config := configdir.LocalConfig("wordpress-cli/config")
	if config != expect {
		t.Fatalf("Expected %s, got %s", expect, config)
	}
}
