package de

import (
	"os"
	"os/exec"
)

type Sway struct{}

func (Sway) SetPape(s string) error {
	swaysock := os.Getenv("SWAYSOCK")
	cmd := exec.Command("swaymsg", "-s", swaysock, "output", "*", "background", s, "fill")
	return cmd.Run()
}
