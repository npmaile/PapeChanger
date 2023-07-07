//go:build linux
// +build linux

package papesetter

import (
	"errors"
	"log"
	"os"
	"os/exec"

	"github.com/cloudflare/ahocorasick"
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
	}
	nextAttempt, err := checkViaPS()
	if err != nil {
		log.Printf("something went wrong with ps to check for your de: %s", err)
		return deNotFound{}
	}
	switch nextAttempt {
	case "sway":
		return de.Sway{}
	}

	return deNotFound{}

}

func checkViaPS() (string, error) {
	ps := exec.Command("ps", "-a")
	output, err := ps.Output()
	if err != nil {
		return "", errors.New("it didn't work")
	}
	matched := ahocorasick.NewMatcher(hardOnes).Match(output)
	if len(matched) == 0 {
		return "", errors.New("desktop environment not found")
	}
	return string(hardOnes[matched[0]]), err
}

var hardOnes = [][]byte{
	[]byte("sway"),
}
