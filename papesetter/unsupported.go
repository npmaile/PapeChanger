//go:build !linux && !darwin
// +build !linux,!darwin

package papesetter

import "fmt"

func SetPape(s string) error {
	return fmt.Errorf("unsupported os")
}
