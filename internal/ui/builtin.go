package ui

import (
	"errors"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func BuiltIn(directories []string, existingApp fyne.App) (string, error) {
	var daemonMode bool
	if existingApp == nil {
		daemonMode = false
		existingApp = app.New()
	}

	var window fyne.Window
	if len(existingApp.Driver().AllWindows()) > 1 {
		fmt.Println(len(existingApp.Driver().AllWindows()))
		window = existingApp.Driver().AllWindows()[1]
	} else {
		window = existingApp.NewWindow("Chooser Widget")
	}
	cont := container.New(layout.NewVBoxLayout())
	selectionChan := make(chan string)
	var closeAction func()
	if daemonMode {
		closeAction = window.Hide
	} else {
		closeAction = window.Close
	}
	for _, item := range directories {
		var item = item
		listItem := widget.NewButton(item, func() {
			go func(item string) {
				selectionChan <- item
			}(item)
			closeAction()
		})
		listItem.Show()
		cont.Add(listItem)
	}
	window.SetContent(container.NewScroll(cont))
	window.Resize(fyne.NewSize(600, 400))
	window.CenterOnScreen()
	window.Show()
	window.RequestFocus()

	oofChan := make(chan struct{})
	if !daemonMode {
		existingApp.Run()
	} else {
		window.SetCloseIntercept(func() {
			window.Hide()
			oofChan <- struct{}{}
		})
	}
	select {
	case ret := <-selectionChan:
		return ret, nil
	case <-oofChan:
		return "", errors.New("no directory picked")
	}
}
