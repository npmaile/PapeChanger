package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/npmaile/papeChanger/internal/environment"
	"github.com/npmaile/papeChanger/internal/errprefix"
	"github.com/npmaile/papeChanger/internal/ui"
	"github.com/npmaile/papeChanger/pkg/papesetter"
)

func main() {
	//changeDir := flag.Bool("c", false, "Change the directory you are selecting walpapers from")
	daemon := flag.Bool("d", false, "run in daemon mode with a status bar icon")
	setup := flag.Bool("setup", false, "set walpaper for the first time")
	flag.Parse()

	env, err := environment.Initialize()
	if err != nil {
		log.Fatalf("%sUnable to initialize environment: %v", errprefix.Get(), err)
	}

	if *setup && !*daemon {
		filepathraw := os.Args[len(os.Args)-1]
		var papePath string
		papePath, err := filepath.Abs(filepathraw)
		if err != nil {
			log.Fatalf("%sUnable to find file %s: %v", errprefix.Get(), filepathraw, err)
		}
		log.Printf("Setting wallpaper to %s", papePath)
		err = papesetter.SetPape(papePath)
		if err != nil {
			log.Fatalf("%sUnable to set walpaper to %s: %v", errprefix.Get(), filepathraw, err)
		}
		err = env.WriteState(papePath)
		if err != nil {
			log.Fatalf("%sUnable to write state file %s: %v", errprefix.Get(), papePath, err)
		}
		os.Exit(0)
	}

	if *daemon {
		ui.RunDaemon(env, *setup)
		os.Exit(0)
	}
}
