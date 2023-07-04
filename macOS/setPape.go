package macos

import (
	"fmt"
	"os/exec"
)

func SetPape(s string) error {
	cmd := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"Finder\" to set desktop picture to POSIX file \"%s\"", s))
	return cmd.Run()
}
