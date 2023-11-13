//go:build darwin
// +build darwin

package papesetter

import (
	"fmt"
	"os/exec"
)

func SetPape(s string) error {
cmd := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"System Events\" to set picture of every desktop to \"%s\"", s))
	return cmd.Run()
}
