package environment

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

type Env struct {
	//runtime config options
	CurrentPape string

	//global config options
	StatePath string
}

func Initialize() (*Env, error) {
	var statePath string
	u, err := user.Current()
	if err != nil {
		return nil, err
	}
	switch runtime.GOOS {
	case "windows":
		statePath = filepath.Join(u.HomeDir, "AppData", "local", "papeChanger", "state")
	default:
		statePath = filepath.Join(u.HomeDir, ".local", "papeChanger", "state")
	}

	currentPapeRaw, err := os.ReadFile(statePath)
	if err != nil {
		return nil, err
	}

	return &Env{
		StatePath:   statePath,
		CurrentPape: string(currentPapeRaw),
	}, nil
}

func (e *Env) WriteState(papePath string) error {
	e.CurrentPape = papePath
	f, err := os.Create(e.StatePath)
	if err != nil {
		return err
	}
	_, err = f.WriteString(papePath)
	return err
}

func (e *Env) PapeDir() string {
	return filepath.Dir(e.CurrentPape)
}
