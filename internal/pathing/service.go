package pathing

import (
	"log"
	"os"
)

// Ensure directories exist on startup
func init() {
	// Directories that must exist:
	dirs := []string{
		GetDataDir(),
	}

	// Create all directories
	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func GetDataDir() string {
	return "/var/lib/home-control-center"
}

func GetConfigDir() string {
	return "/etc/home-control-center"
}
