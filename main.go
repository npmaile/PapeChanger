package main

import (
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
	"github.com/pborman/getopt"
	"golang.org/x/exp/slog"
)

func main() {
	setupMode := false
	changeDir := getopt.BoolLong("change",    'c', "change wallpaper subdirectory in rofi")
	previous  := getopt.BoolLong("previous",  'p', "restore to last set wallpaper (useful to run on startup)")
	daemon    := getopt.BoolLong("daemon",     0,  "run in daemon mode with a status bar icon")
	help      := getopt.BoolLong("help",       0,  "print this page and exit")
	getMaster := getopt.BoolLong("getmaster", 'g', "print wallpaper master directory path")
	randomize := getopt.BoolLong("randomize", 'r', "select random wallpaper instead of choosing one")

	selectDir := getopt.StringLong("directory", 'd', "", "manually set wallpaper directory and choose one in it")
	setup     := getopt.StringLong("setup",     's', "",  "accepts a path to a wallpaper in the master directory\n\n"+
		"Example master directory structure:\n\n"+
		"wallpapers\n"+
		"├── nature\n"+
		"│   ├── eagle.jpg\n"+
		"│   ├── bear.jpg\n"+
		"│   └── deer.jpg\n"+
		"└── cars\n"+
		"    ├── lambo.jpg\n"+
		"    ├── gtr.jpg\n"+
		"    └── wrx.jpg\n")

	getopt.Parse()

	if *help {
		getopt.Usage()
		os.Exit(0)
	}
	

	var env *environment.Env
	var err error


	if *previous {
		env, err = environment.GetState()
		if err != nil {
			fatalf("%sUnable to get state of papechanger environment: %v", errprefix.Get(), err)
		}
		err = papesetter.SetPape(env.CurrentPape)
		if err != nil {
			fatalf("%sUnable to set wallpaper to %s: %v", errprefix.Get(), env.CurrentPape, err)
		}
		return
	}

	if *getMaster {
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
			fatalf("%sUnable to select background in \"%s\"  %v", errprefix.Get(), *selectDir, err)
		}
		err = papesetter.SetPape(pape)
		if err != nil {
			fatalf("%sUnable to set wallpaper: %v", errprefix.Get(), err)
		}
		env.WriteState(pape)

		return
	}

	if setup != nil && *setup != "" && !*daemon {
		setupMode = true
		var papePath string
		papePath, err = filepath.Abs(*setup)

		if err != nil {
			fatalf("%sUnable to find path %s: %v", errprefix.Get(), *setup, err)
		}

		log.Printf("Setting wallpaper path to %s", papePath)
		err = papesetter.SetPape(papePath)

		if err != nil {
			fatalf("%sUnable to set walpaper to %s: %v", errprefix.Get(), *setup, err)
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
		ui.RunDaemon(env, setupMode)
	} else {
		classicFunctionality(env, *randomize, *changeDir)
	}
}

func classicFunctionality(env *environment.Env, randomize bool, changeDir bool) {
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
	} else {
		pape2Pick, err = selector.SelectWallpaperInOrder(papeDir, env.CurrentPape)
	}

	if err != nil {
		fatalf("%sUnable to select Wallpaper \"%s\": %v", errprefix.Get(), pape2Pick, err)
	}
	log.Printf("Setting wallpaper to %s", pape2Pick)
	err = papesetter.SetPape(pape2Pick)
	if err != nil {
		fatalf("%sUnable to set \"%s\" as wallpaper: %v", errprefix.Get(), pape2Pick, err)
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
