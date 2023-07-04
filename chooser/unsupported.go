//go:build !linux && !darwin
// +build !linux,!darwin

package chooser

import "fmt"

func Chooser([]string) (string, error) {
	return "", fmt.Errorf("Unsupported OS")
}
