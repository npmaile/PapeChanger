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

/*
func RunDaemon(doWorkfunction func()) {
	switch runtime.GOOS {
	case "linux", "darwin":
		papeChangerLogo = iconPng
		xIcon = xIconPng
	case "windows":
		papeChangerLogo = iconIco
		xIcon = xIconIco
	default:
		log.Fatal("you've done something horribly wrong")
	}

	systray.SetIcon(papeChangerLogo)
	systray.SetTitle("PapeChanger")
	systray.SetTooltip("This is the papechanger systray icon you receive when running the application in daemon mode.")
	mOpen := systray.AddMenuItem("Change Directory", "change the directory your application is looking for shit in cuh")
	mOpen.SetIcon(papeChangerLogo)
	mQuit := systray.AddMenuItem("Quit", "I'm not sure why you'd ever wish to quit this perfect application, but if you're stupid and hate aesthics, feel free.")
	mQuit.SetIcon(xIcon)
	go func() {
		select {
		case <-mQuit.ClickedCh:
			systray.Quit()
		case <-mOpen.ClickedCh:
			doWorkfunction()
		}
	}()
	fmt.Println("systray started")
	systray.Run(func() {}, func() {})
	fmt.Println("systray ended")
}
*/

func RunDaemon(doWorkFunction func(bool)) {
	a := app.New()
	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("PapeChanger",
			fyne.NewMenuItem("Change Wallpaper", func() {
				doWorkFunction(false)
			}),
			fyne.NewMenuItem("Change Directory", func() {
				doWorkFunction(true)
			}),
		)
		desk.SetSystemTrayMenu(m)
	}
	a.Run()
}
