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
	setup := flag.Bool("setup", false, "set walpaper for the first time")
	flag.Parse()


	if *setup && !*daemon {
		filepathraw := os.Args[len(os.Args)-1]
		var papePath string
		papePath, err := filepath.Abs(filepathraw)
		if err != nil {
			fatalf("%sUnable to find file %s: %v", errprefix.Get(), filepathraw, err)
		}
		log.Printf("Setting wallpaper to %s", papePath)
		err = papesetter.SetPape(papePath)
		if err != nil {
			fatalf("%sUnable to set walpaper to %s: %v", errprefix.Get(), filepathraw, err)
		}
		err = env.WriteState(papePath)
		if err != nil {
			fatalf("%sUnable to write state file %s: %v", errprefix.Get(), papePath, err)
		}
		os.Exit(0)
	}
	
	env, err := environment.Initialize()
	if err != nil {
		fatalf("%sUnable to initialize environment: %v", errprefix.Get(), err)
	}

	if *daemon {
		fmt.Println(env.DirOfDirs())
		ui.RunDaemon(env, *setup)
	} else {
		classicFunctionality(env, *changeDir)
	}
}

func classicFunctionality(env *environment.Env, changeDir bool) {
	var papeDir string
	if changeDir {
		dirs, err := selector.ListDirectories(env.DirOfDirs())
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
	pape2Pick, err := selector.SelectWallpaper(papeDir)
	if err != nil {
		fatalf("%sUnable to select Wallpaper: %v", errprefix.Get(), err)
	}
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
