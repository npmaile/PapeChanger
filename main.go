package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/npmaile/papeChanger/internal/chooser"
	"github.com/npmaile/papeChanger/internal/environment"
	"github.com/npmaile/papeChanger/internal/errprefix"
	"github.com/npmaile/papeChanger/internal/selector"
	"github.com/npmaile/papeChanger/internal/ui"
	"github.com/npmaile/papeChanger/pkg/papesetter"
	"golang.org/x/exp/slog"
)

func main() {
	changeDir := flag.Bool("c", false, "Change the directory you are selecting walpapers from")
	daemon := flag.Bool("d", false, "run in daemon mode with a status bar icon")
	listDirOfDirs := flag.Bool("papeDirsDir", false, "interrogate the command line application to determine the directory containing all wallpaper directories")
	selectDir := flag.String("directory", "", "manually set the directory to be used by papechanger and change the wallpaper to one in it")
	setup := flag.Bool("setup", false, "-setup <path> provide a path to the directory with image files\n\n"+
		"Typically you would set this up like this:\n\n"+
		"wallpapers\n"+
		"├── nature\n"+
		"│   ├── eagle.jpg\n"+
		"│   ├── bear.jpg\n"+
		"│   └── deer.jpg\n"+
		"└── cars\n"+
		"    ├── lambo.jpg\n"+
		"    ├── gtr.jpg\n"+
		"    └── wrx.jpg\n")
	randomize := flag.Bool("randomize", false, "select a random wallpaper instead of going through the list of wallpapers directly")
    papeDownload := flag.Bool("url", false, "-url <URL> provide a valid URL to the image you wish to download\n")
	flag.Parse()

	var env *environment.Env
	var err error

	if *listDirOfDirs {
		env, err = environment.GetState()
		if err != nil {
			fatalf("%sUnable to get state of papechanger environment: %v", errprefix.Get(), err)
		}
		fmt.Println(env.DirOfDirs())
		return
	}

	if selectDir != nil && *selectDir != "" {
		env, err = environment.GetState()
		if err != nil {
			fatalf("%sUnable to get state of papechanger environment: %v", errprefix.Get(), err)
		}
		var pape string
		pape, err = selector.SelectWallpaperRandom(*selectDir)
		if err != nil {
			fatalf("%sUnable to select a wallpaper: %v", errprefix.Get(), err)
		}
		err = papesetter.SetPape(pape)
		if err != nil {
			fatalf("%sUnalbe to set wallpaper: %v", errprefix.Get(), err)
		}
		env.WriteState(pape)

		return
	}

	if *setup && !*daemon {

		filepathraw := os.Args[len(os.Args)-1]
		var papePath string
		papePath, err = filepath.Abs(filepathraw)

		if err != nil {
			fatalf("%sUnable to find path %s: %v", errprefix.Get(), filepathraw, err)
		}

		log.Printf("Setting wallpaper path to %s", papePath)
		err = papesetter.SetPape(papePath)

		if err != nil {
			fatalf("%sUnable to set walpaper to %s: %v", errprefix.Get(), filepathraw, err)
		}

		_, err = environment.InitializeState(papePath)

		if err != nil {
			fatalf("%sUnable to write state file %s: %v", errprefix.Get(), papePath, err)
		}

		os.Exit(0)

	} else {

		env, err = environment.GetState()

	}

	if err != nil {
		fatalf("%sUnable to initialize environment: %v", errprefix.Get(), err)
	}

	if *daemon {
		ui.RunDaemon(env, *setup)
	} else {
		classicFunctionality(env, *randomize, *changeDir, *papeDownload)
	}
}

func classicFunctionality(env *environment.Env, randomize bool, changeDir bool, papeDownload bool) {
	var papeDir string
	if changeDir {

		dirs, err := selector.ListDirectories(env.DirOfDirs())

		log.Printf("Found %d directories", len(dirs))

		if err != nil {
			fatalf("%sUnable to change wallpaper directory: %v", errprefix.Get(), err)
		}
		dirToPick, err := chooser.Chooser(dirs)
		if err != nil {
			fatalf("%sUnable to pick wallpaper directory: %v", errprefix.Get(), err)
		}
		papeDir = fmt.Sprintf("%s%s%s", env.DirOfDirs(), string(os.PathSeparator), dirToPick)

	}
	if papeDir == "" {
		papeDir = env.PapeDir()
	}

	var pape2Pick string
	var err error
    if randomize {
        pape2Pick, err = selector.SelectWallpaperRandom(papeDir)
    } else if papeDownload {
        url := os.Args[len(os.Args)-1]
        pape2Pick, err = selector.SelectDownloadedWallpaper(papeDir, url)
    } else {
        pape2Pick, err = selector.SelectWallpaperInOrder(papeDir, env.CurrentPape)
    }

	if err != nil {
		fatalf("%sUnable to select Wallpaper: %v", errprefix.Get(), err)
	}
	log.Printf("Setting wallpaper to %s", pape2Pick)
	err = papesetter.SetPape(pape2Pick)
	if err != nil {
		fatalf("%sUnable to set wallpaper: %v", errprefix.Get(), err)
	}
	err = env.WriteState(pape2Pick)
	if err != nil {
		fatalf("%sUnable to write state file: %v", errprefix.Get(), err)
	}
}

func fatalf(msg string, vars ...any) {
	slog.Error(fmt.Sprintf(msg, vars...))
	os.Exit(1)
}

