//go:build !linux && !solaris
// +build !linux,!solaris

package files

import (
	"errors"
)

// copy directories recursively from source path to destination path
func copyDirectory(srcDir, dest string) error {
	return errors.New("copy directories is available only on linux")
}
