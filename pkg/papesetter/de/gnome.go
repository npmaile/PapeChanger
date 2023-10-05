package de

import (
	"fmt"
	"os/exec"
)

type Gnome struct{}

func (Gnome) SetPape(s string) error {
	cmd := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", s)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Unable to set wallpaper: %s", err)
	}
	return nil
}
