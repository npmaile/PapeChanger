package fileresolver

import (
	"flag"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"github.com/npmaile/papeChanger/internal/chooser"
	"github.com/npmaile/papeChanger/internal/errprefix"
	"github.com/npmaile/papeChanger/internal/ui"
	"github.com/npmaile/papeChanger/pkg/papesetter"
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func main() {
	useBuiltInChanger := flag.Bool("useBuiltin", false, "Use the built-in selector widget instead of one you have installed")
	changeDir := flag.Bool("c", false, "Change the directory you are selecting walpapers from")
	randomize := flag.Bool("r", true, "Randomize wallpaper to change")
	daemon := flag.Bool("d", false, "run in daemon mode with a status bar icon")
	setup := flag.Bool("setup", false, "set walpaper for the first time")
	u, err := user.Current()
	if err != nil {
		log.Fatalf("%sHow the H**K are you not logged in as a user?", errprefix.Get())
	}
	homeDir := u.HomeDir
	var stateFile *string
	switch runtime.GOOS {
	case "windows":
		stateFile = flag.String("m", filepath.Join(homeDir, "AppData", "Local", "papeChanger", "state"), "Use a custom location to store the current walpaper set")
	default:
		stateFile = flag.String("m", filepath.Join(homeDir, ".local", "papeChanger", "state"), "Use a custom location to store the current walpaper set")
	}
	flag.Parse()

	if *setup {
		filepathraw := os.Args[len(os.Args)-1]
		var papePath string
		papePath, err = filepath.Abs(filepathraw)
		if err != nil {
			log.Fatalf("%sUnable to find file %s: %v", errprefix.Get(), filepathraw, err)
		}
		log.Printf("Setting wallpaper to %s", papePath)
		err = papesetter.SetPape(papePath)
		if err != nil {
			log.Fatalf("%sUnable to set walpaper to %s: %v", errprefix.Get(), filepathraw, err)
		}
		err = writeState(*stateFile, papePath)
		if err != nil {
			log.Fatalf("%sUnable to write state file %s: %v", errprefix.Get(), *stateFile, err)
		}
		os.Exit(0)
	}

	if *daemon {
		ui.RunDaemon(func(changeDir bool, existingApp fyne.App) {
			t := true
			doWork(&t, &changeDir, randomize, stateFile, existingApp)
		})
	}

	doWork(useBuiltInChanger, changeDir, randomize, stateFile, nil)
}

func writeState(stateFile string, newWalpaper string) error {
	f, err := os.Create(stateFile)
	if err != nil {
		return err
	}
	_, err = f.Write([]byte(newWalpaper))
	return err
}

func doWork(useBuiltInChanger *bool, changeDir *bool, randomize *bool, stateFile *string, existingApp fyne.App) {
	currentWalpaper, err := os.ReadFile(*stateFile)
	if err != nil {
		log.Fatalf("%sCan't read the file: %v", errprefix.Get(), err)
	}
	var pathParts []string
	switch runtime.GOOS {
	case "windows":
		pathParts = strings.Split(string(currentWalpaper), string(os.PathSeparator))
	default:
		pathParts = strings.Split(string(currentWalpaper), string(os.PathSeparator))
	}
	currentDirParts := pathParts[0 : len(pathParts)-1]
	if *changeDir {
		var megaDir string
		switch runtime.GOOS {
		case "windows":
			folderParts := append([]string{currentDirParts[0], "\\"}, currentDirParts[1:len(currentDirParts)-1]...)
			megaDir = string(filepath.Join(folderParts...))

		default:
			megaDir = string(os.PathSeparator) + filepath.Join(currentDirParts[0:len(currentDirParts)-1]...)
		}
		var files []fs.DirEntry
		files, err = os.ReadDir(megaDir)
		if err != nil {
			log.Fatalf("%sYou've moved your walpapers around and I can't find them now: %v", errprefix.Get(), err)
		}
		dirList := []string{}
		for _, file := range files {
			if file.IsDir() {
				dirList = append(dirList, file.Name())
			}
		}
		var chosen string
		if !*useBuiltInChanger {
			chosen, err = chooser.Chooser(dirList)
			if err != nil {
				log.Fatalf("%sFailed to choose walpaper directory: %v", errprefix.Get(), err)
			}
		} else {
			chosen, err = chooser.BuiltIn(dirList, existingApp)
			if err != nil {
				log.Fatalf("%sFailed to choose walpaper directory: %v", errprefix.Get(), err)
			}
		}
		currentDirParts[len(currentDirParts)-1] = string(chosen)
	}

	var walpaperFolder string
	switch runtime.GOOS {
	case "windows":
		folderParts := append([]string{currentDirParts[0], "\\"}, currentDirParts[1:]...)
		walpaperFolder = string(filepath.Join(folderParts...))
	default:
		walpaperFolder = string(os.PathSeparator) + filepath.Join(currentDirParts...)
	}
	papers, err := os.ReadDir(walpaperFolder)
	if err != nil {
		log.Fatalf("%sUnable to get list of individual walpapers: %v", errprefix.Get(), err)
	}

	var fullPath []string
	if *randomize {
		index := rand.Int() % len(papers)
		fullPath = append(currentDirParts, papers[index].Name())
	} else {
		//todo
	}
	var newWalpaper string
	switch runtime.GOOS {
	case "windows":
		folderParts := append([]string{fullPath[0], "\\"}, fullPath[1:]...)
		newWalpaper = string(filepath.Join(folderParts...))
	default:
		newWalpaper = string(os.PathSeparator) + filepath.Join(fullPath...)
	}
	err = papesetter.SetPape(newWalpaper)
	if err != nil {
		log.Printf("%sunable to change walpaper: %v", errprefix.Get(), err)
	}
	err = writeState(*stateFile, newWalpaper)
	if err != nil {
		log.Printf("%sCreation of state file failed: %v", errprefix.Get(), err)
	}
}
