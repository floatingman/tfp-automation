package rancher2

import (
	"fmt"
	"os"
	"path/filepath"
)

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Keep going up until we find go.mod
	for {
		// Check if go.mod exists in current directory
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		// Get the parent directory
		parent := filepath.Dir(dir)

		// If we're already at the root and haven't found go.mod, stop
		if parent == dir {
			return "", fmt.Errorf("could not find project root (no go.mod file found)")
		}

		dir = parent
	}
}

// SetKeyPath is a function that will set the path to the key file.
func SetKeyPath(keyPath string) string {
	userDir, err := findProjectRoot()
	if err != nil {
		return ""
	}

	keyPath = filepath.Join(userDir, keyPath)
	return keyPath
}
