package ui

import (
	_ "embed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
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
