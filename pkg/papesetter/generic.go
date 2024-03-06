package papesetter

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
)

var shell string

func init() {
	switch runtime.GOOS {
	case "windows":
		shell = "powershell"
	case "linux", "darwin":
		shell = "/bin/sh"

	default:
		fmt.Println("todo: a message about operating systems, but we will give it a shot anyway")
		shell = "/bin/sh"

	}
}

func SetPapeCustom(papePath string, command string) error {
	cmd := exec.Command(shell, fmt.Sprintf(command, papePath))
	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("todo")
	}
	errorPipe, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("todo")
	}

	go func() {
		_, _ = io.Copy(os.Stdout, outPipe)
		_, _ = io.Copy(os.Stderr, errorPipe)
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Println("todo")
	}
	err = cmd.Wait()

	return err
}
