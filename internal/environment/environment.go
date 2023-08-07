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

func InitializeState(firstPape string) (*Env, error) {
	statePath, err := StatePath()
	if err != nil{
		return nil, err
	}
	e := Env{StatePath:statePath,CurrentPape:firstPape}
	err = e.WriteState(firstPape)
	if err != nil{
		return nil, err
	}
	return GetState()
}

func GetState() (*Env, error) {
	statePath, err := StatePath()
	if err != nil {
		return nil, err
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
	statePath, err := StatePath()
	if err != nil {
		return err
	}
	statePath = filepath.Dir(statePath)
	err = os.MkdirAll(statePath, 0777)
	if err != nil {
		return err
	}
	f, err := os.Create(e.StatePath)
	if err != nil {
		return err
	}
	_, err = f.WriteString(papePath)
	return err
}

func StatePath() (string, error) {
	var statePath string
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	switch runtime.GOOS {
	case "windows":
		statePath = filepath.Join(u.HomeDir, "AppData", "local", "papeChanger", "state")
	default:
		statePath = filepath.Join(u.HomeDir, ".local", "papeChanger", "state")
	}
	return statePath, nil

}

func (e *Env) PapeDir() string {
	return filepath.Dir(e.CurrentPape)
}

func (e *Env) DirOfDirs() string {
	return filepath.Dir(e.PapeDir())
}
