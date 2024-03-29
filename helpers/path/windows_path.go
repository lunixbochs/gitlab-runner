// +build windows

// This implementation only works when compiled for Windows
// as this uses the `path/filepath` which is platform dependent

package path

import (
	"path/filepath"
	"strings"
)

type windowsPath struct{}

func (p *windowsPath) Join(elem ...string) string {
	return filepath.Join(elem...)
}

func (p *windowsPath) IsAbs(path string) bool {
	path = filepath.Clean(path)
	return filepath.IsAbs(path)
}

func (p *windowsPath) IsRoot(path string) bool {
	path = filepath.Clean(path)
	return filepath.IsAbs(path) && filepath.Dir(path) == path
}

func (p *windowsPath) Contains(basePath, targetPath string) bool {
	// we use `filepath.Rel` as this perform OS-specific comparision
	// and this set of functions is compiled using OS-specific golang filepath
	relativePath, err := filepath.Rel(basePath, targetPath)
	if err != nil {
		return false
	}

	// if it starts with `..` it tries to escape the path
	if strings.HasPrefix(relativePath, "..") {
		return false
	}

	return true
}

func NewWindowsPath() Path {
	return &windowsPath{}
}
