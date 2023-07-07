//go:build linux
// +build linux

package papesetter

import (
	"os"

	"github.com/npmaile/papeChanger/papesetter/de"
)

func SetPape(s string) error {
	desktop := getDesktop()
	return desktop.SetPape(s)
}

func getDesktop() DE {
	switch os.Getenv("XDG_CURRENT_DESKTOP") {
	case "KDE":
		return de.Plasma{}
	default: // pass through to next
	}
	return deNotFound{}

}
