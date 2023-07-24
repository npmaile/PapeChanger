package ui

import (
	_ "embed"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/npmaile/papeChanger/internal/environment"
	"github.com/npmaile/papeChanger/internal/selector"
	"github.com/npmaile/papeChanger/pkg/papesetter"
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

func RunDaemon(env *environment.Env, setup bool) {
	a := app.New()
	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("PapeChanger",
			fyne.NewMenuItem("Change Wallpaper", func() {
				nextPape, err := selector.SelectWallpaper(env.PapeDir())
				if err != nil {
					log.Printf("unable to change wallpaper: %s", err.Error())
				}
				papesetter.SetPape(nextPape)
				env.WriteState(nextPape)

			}),
			fyne.NewMenuItem("Change Directory", func() {
				dirs, err := selector.ListDirectories(env.PapeDir())
				if err != nil {
					log.Printf("unable to change directory :%s", err.Error())
				}
				selectionChan := make(chan string, 1)
				ChooserWindow(a, dirs, selectionChan)
				selection := <-selectionChan
				nextPape, err := selector.SelectWallpaper(selection)
				if err != nil {
					log.Printf("unable to change wallpaper: %s", err.Error())
				}
				papesetter.SetPape(nextPape)
				env.WriteState(nextPape)
			}),
		)
		desk.SetSystemTrayMenu(m)
	}

	if setup {
		selection := make(chan string, 1)
		SetupWindow(a, selection)
		s := <-selection
		env.WriteState(s)
		papesetter.SetPape(s)
	}
	a.Run()
}

func SetupWindow(app fyne.App, selectedPapePath chan string) {
	window := app.NewWindow("Select a wallpaper")
	dialog.NewFileOpen(
		func(rc fyne.URIReadCloser, e error) {
			selectedPapePath <- rc.URI().Path()
		},
		window)
	window.Show()
}

func ChooserWindow(app fyne.App, directories []string, selectionChan chan string) {
	window := app.NewWindow("select directory")
	cont := container.New(layout.NewVBoxLayout())
	for _, item := range directories {
		var item = item
		listItem := widget.NewButton(item, func() {
			go func(item string) {
				selectionChan <- item
			}(item)
			window.Hide()
		})
		listItem.Show()
		cont.Add(listItem)
	}
	window.SetContent(container.NewScroll(cont))
	window.Resize(fyne.NewSize(600, 400))
	window.CenterOnScreen()
}
