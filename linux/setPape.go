package linux

import (
	"os"
	"os/exec"
)

func SetPape(s string) error {
	swaysock := os.Getenv("SWAYSOCK")
	cmd := exec.Command("swaymsg", "-s", swaysock, "output", "*", "background", s, "fill")
	return cmd.Run()
}
