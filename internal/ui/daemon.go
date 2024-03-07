package ui

import (
	_ "embed"
	"errors"
	"fmt"
	"log"
	"os"

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
				nextPape, err := selector.SelectWallpaperRandom(env.PapeDir())
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
				selectionChan := make(chan StringSelectionWithErr, 1)
				ChooserWindow(a, dirs, selectionChan)
				selection := <-selectionChan
				if selection.Err != nil {
					log.Printf("unable to retrieve directory selection: %s", selection.Err.Error())
					return
				}

				selectionFullPath := fmt.Sprintf("%s%s%s", env.DirOfDirs(), string(os.PathSeparator), selection.SelectedItem)
				nextPape, err := selector.SelectWallpaperRandom(selectionFullPath)
				if err != nil {
					log.Printf("unable to change wallpaper: %s", err.Error())
					return
				}
				err = papesetter.SetPape(nextPape)
				if err != nil {
					log.Printf("unable to set wallpaper: %s", err.Error())
					return
				}
				err = env.WriteState(nextPape)
				if err != nil {
					log.Printf("unable to write state: %s", err.Error())
					return
				}
			}),
		)
		desk.SetSystemTrayMenu(m)
		desk.SetSystemTrayIcon(fyne.NewStaticResource("icon", iconPng))
	}

	if setup {
		selection := make(chan StringSelectionWithErr, 1)
		SetupWindow(a, selection)
		s := <-selection
		if s.Err != nil {
			//todo: something
		}
		env.WriteState(s.SelectedItem)
		papesetter.SetPape(s.SelectedItem)
	}
	a.Run()
}

type StringSelectionWithErr struct {
	Err          error
	SelectedItem string
}

func SetupWindow(app fyne.App, selectedPapePath chan StringSelectionWithErr) {
	window := app.NewWindow("Select a wallpaper")
	dialog.NewFileOpen(
		func(rc fyne.URIReadCloser, e error) {
			if e != nil {
				selectedPapePath <- StringSelectionWithErr{Err: e, SelectedItem: ""}
			}
			selectedPapePath <- StringSelectionWithErr{SelectedItem: rc.URI().Path()}

		},
		window)
	window.Show()
}

func ChooserWindow(app fyne.App, directories []string, selectionChan chan StringSelectionWithErr) {
	fmt.Printf("length passed to chooser window: %d\n", len(directories))
	window := app.NewWindow("select directory")
	cont := container.New(layout.NewVBoxLayout())
	for _, item := range directories {
		var item = item
		listItem := widget.NewButton(item, func() {
			go func(item string) {
				selectionChan <- StringSelectionWithErr{SelectedItem: item}
			}(item)
			window.Hide()
			window.Close()
		})
		cont.Add(listItem)
	}
	window.SetCloseIntercept(func() {
		selectionChan <- StringSelectionWithErr{Err: errors.New("window closed without selection")}
		window.Hide()
		window.Close()
	})
	fmt.Println("here")

	window.SetContent(container.NewScroll(cont))
	window.Resize(fyne.NewSize(600, 400))
	window.CenterOnScreen()
	window.Show()
	fmt.Println("done with the show")
}
