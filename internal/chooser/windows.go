//go:build windows
// +build windows

package chooser

import (
	"fyne.io/fyne/v2/app"
	"github.com/npmaile/papeChanger/internal/ui"
)

func Chooser(dirs []string) (string, error) {
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
