package chooser

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"

	"fyne.io/fyne/v2/app"
	"github.com/npmaile/papeChanger/internal/ui"
)

func BuiltIn(dirs []string) (string, error) {
	// will replace with something else if I find a good chooser for windows
	app := app.New()
	selectionChan := make(chan ui.StringSelectionWithErr)
	ui.ChooserWindow(app, dirs, selectionChan)
	selection := <-selectionChan
	if selection.Err != nil {
		return "", selection.Err
	}
	return selection.SelectedItem, nil
}

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

func UserDefined(dirs []string, command string) (string, error) {
	cmd := exec.Command(shell, "-c", command)
	inPipe, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("todo")
	}

	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("todo")
	}

	errorPipe, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("todo")
	}

	go func() {
		for _, dir := range dirs {
			_, err = io.WriteString(inPipe, dir+"\n")
			if err != nil {
				fmt.Println("todo")
			}
			inPipe.Close()
		}
	}()

	outputBuffer := bytes.NewBuffer([]byte{})
	go func() {
		_, _ = io.Copy(outputBuffer, outPipe)
		_, _ = io.Copy(os.Stderr, errorPipe)
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Println("todo")
	}
	err = cmd.Wait()

	return outputBuffer.String(), nil
}
