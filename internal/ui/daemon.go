package ui

import (
	_ "embed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

//go:embed icons/icon.png
var iconPng []byte

//go:embed icons/x.png
var xIconPng []byte

//go:embed icons/icon.ico
var iconIco []byte

//go:embed icons/x.ico
var xIconIco []byte

var papeChangerLogo []byte
var xIcon []byte

func RunDaemon(doWorkFunction func(bool, fyne.App)) {
	a := app.New()
	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("PapeChanger",
			fyne.NewMenuItem("Change Wallpaper", func() {
				doWorkFunction(false, a)
			}),
			fyne.NewMenuItem("Change Directory", func() {
				doWorkFunction(true, a)
			}),
		)
		desk.SetSystemTrayMenu(m)
	}
	a.Run()
}

func ChooserWindow(directories []string) fyne.Window {
	var window fyne.Window
	cont := container.New(layout.NewVBoxLayout())
	selectionChan := make(chan string)
	var closeAction func()
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
	return window
}
