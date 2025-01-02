package configdir

import "path/filepath"

func init() {
	findPaths()
}

func LocalConfig(folder ...string) string {
	if len(folder) == 0 {
		return localConfig
	}

	return filepath.Join(localConfig, filepath.Join(folder...))
}
