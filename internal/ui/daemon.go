package ui

import (
	_ "embed"
	"log"
	"os"
	"fmt"

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
					return
				}
				papesetter.SetPape(nextPape)
				env.WriteState(nextPape)

			}),
			fyne.NewMenuItem("Change Directory", func() {
				dirs, err := selector.ListDirectories(env.DirOfDirs())
				if err != nil {
					log.Printf("unable to change directory :%s", err.Error())
					return
				}
				selectionChan := make(chan string, 1)
				ChooserWindow(a, dirs, selectionChan)
				selection := <-selectionChan

				selectionFullPath := fmt.Sprintf("%s%s%s", env.DirOfDirs(), string(os.PathSeparator), selection)
				nextPape, err := selector.SelectWallpaper(selectionFullPath)
				if err != nil {
					log.Printf("unable to change wallpaper: %s", err.Error())
					return
				}
				err = papesetter.SetPape(nextPape)
				if err != nil{
					log.Printf("unable to set wallpaper: %s", err.Error())
					return
				}
				err = env.WriteState(nextPape)
				if err != nil{
					log.Printf("unable to write state: %s", err.Error())
					return
				}
			}),
		)
		desk.SetSystemTrayMenu(m)
	}

	if setup {
		selection := make(chan papePathSelection, 1)
		SetupWindow(a, selection)
		s := <-selection
		if s.err != nil {
			//todo: something
		}
		env.WriteState(s.selectedPape)
		papesetter.SetPape(s.selectedPape)
	}
	a.Run()
}

type papePathSelection struct {
	err          error
	selectedPape string
}

func SetupWindow(app fyne.App, selectedPapePath chan papePathSelection) {
	window := app.NewWindow("Select a wallpaper")
	dialog.NewFileOpen(
		func(rc fyne.URIReadCloser, e error) {
			if e != nil {
				selectedPapePath <- papePathSelection{err: e, selectedPape: ""}
			}
			selectedPapePath <- papePathSelection{selectedPape: rc.URI().Path()}

		},
		window)
	window.Show()
}

func ChooserWindow(app fyne.App, directories []string, selectionChan chan string) {
	fmt.Printf("lengh passed to chooser window: %d\n", len(directories))
	window := app.NewWindow("select directory")
	cont := container.New(layout.NewVBoxLayout())
	for _, item := range directories {
		var item = item
		listItem := widget.NewButton(item, func() {
			go func(item string) {
				selectionChan <- item
			}(item)
			window.Hide()
			window.Close()
		})
		cont.Add(listItem)
	}
	window.SetContent(container.NewScroll(cont))
	window.Resize(fyne.NewSize(600, 400))
	window.CenterOnScreen()
	window.Show()
}
