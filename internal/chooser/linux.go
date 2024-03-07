//go:build linux
// +build linux

package chooser

import (
	"io"
	"os/exec"
	"strings"
)

func Chooser(directories []string) (string, error) {
	//Todo: come up with a bespoke list of choosers that will work okay for the user to use
	cmd := exec.Command("rofi", "-dmenu")
	var pipe io.WriteCloser
	pipe, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}
	var outPipe io.ReadCloser
	outPipe, err = cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	err = cmd.Start()
	if err != nil {
		return "", err
	}
	pipe.Write([]byte(strings.Join(directories, "\n")))
	pipe.Close()

	var pickedFile []byte
	pickedFile, err = io.ReadAll(outPipe)
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(pickedFile), "\n"), nil
}
