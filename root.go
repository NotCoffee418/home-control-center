package homecontrolcenter

import (
	"embed"
	"io/fs"
)

//go:embed frontend/dist/*
var frontendFS embed.FS

// GetFrontendFS returns the embedded frontend filesystem
func GetFrontendFS() (fs.FS, error) {
	return fs.Sub(frontendFS, "frontend/dist")
}
