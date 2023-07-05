//go:build !linux && !darwin && !windows
// +build !linux,!darwin,!windows

package chooser

import "fmt"

func Chooser([]string) (string, error) {
	return "", fmt.Errorf("Unsupported OS")
}
