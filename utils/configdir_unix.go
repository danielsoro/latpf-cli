package configdir

import (
	"os"
)

var localConfig string

func findPaths() {
	localConfig = os.Getenv("HOME") + "/.config"
}
