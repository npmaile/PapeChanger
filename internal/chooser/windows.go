//go:build windows
// +build windows

package chooser

import(
	"github.com/npmaile/papeChanger/internal/ui"
	"fyne.io/fyne/v2/app"
)

func Chooser(dirs []string) (string, error) {
	// will replace with something else if I find a good chooser for windows
	app := app.New()
	selectionChan := make(chan string)
	ui.ChooserWindow(app, dirs, selectionChan)
	selection := <- selectionChan
	return selection, nil
}
