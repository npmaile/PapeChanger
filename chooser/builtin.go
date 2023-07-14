package chooser

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func BuiltIn(directories []string, existingApp fyne.App) (string, error) {
	var run bool
	if existingApp == nil {
		run = true
		existingApp = app.New()
	}
	window := existingApp.NewWindow("Chooser Widget")
	window.CenterOnScreen()
	cont := container.New(layout.NewVBoxLayout())

	selectionChan := make(chan string)

	for _, item := range directories {
		var item = item
		listItem := widget.NewButton(item, func() {
			go func(item string) {
				selectionChan <- item
			}(item)
			window.Close()
		})
		listItem.Show()
		cont.Add(listItem)
	}

	window.SetContent(container.NewScroll(cont))
	window.Resize(fyne.NewSize(600, 400))
	window.Show()
	if run{
		existingApp.Run()
	}
	s := <-selectionChan
	return s, nil
}

func lockInSelection(selection string, window fyne.Window) string {
	window.Close()
	return selection
}
