//go:build !linux && !darwin && !windows
// +build !linux,!darwin,!windows

package papesetter

import "fmt"

func SetPape(s string) error {
	return fmt.Errorf("unsupported os")
}
