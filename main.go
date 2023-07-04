package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

const errPrefix = "😭😭oOpSy DoOpSiE, you made a frickey-wickey 😭😭: "

func main() {
	// parse command line arguments
	changeDir := flag.Bool("c", false, "Change the directory you are selecting walpapers from")
	randomize := flag.Bool("r", true, "Randomize wallpaper to change")
	u, err := user.Current()
	if err != nil {
		log.Fatalf("%sHow the H**K are you not logged in as a user?", errPrefix)
	}
	homeDir := u.HomeDir
	stateFile := flag.String("m", path.Join(homeDir, ".local/papeChanger/state"), "Use a custom location to store the current walpaper set")
	flag.Parse()

	currentWalpaper, err := os.ReadFile(*stateFile)
	if err != nil {
		log.Fatalf("%sCan't read the file", errPrefix)
	}
	pathParts := strings.Split(string(currentWalpaper), string(os.PathSeparator))
	currentDirParts := pathParts[0 : len(pathParts)-1]
	if *changeDir {
		megaDir := string(os.PathSeparator) + filepath.Join(currentDirParts[0:len(currentDirParts)-1]...)
		var files []fs.DirEntry
		files, err = os.ReadDir(megaDir)
		if err != nil {
			log.Fatalf("%sYou've moved your walpapers around and I can't find them now: %e", errPrefix, err)
		}
		dirList := []string{}
		for _, file := range files {
			if file.IsDir() {
				dirList = append(dirList, file.Name())
			}
		}
		cmd := exec.Command("choose")
		var pipe io.WriteCloser
		pipe, err = cmd.StdinPipe()
		if err != nil {
			log.Fatalf("%sCan't connect to standard in pipe of chooser: %e", errPrefix, err)
		}
		var outPipe io.ReadCloser
		outPipe, err = cmd.StdoutPipe()
		if err != nil {
			log.Fatalf("%sFailed to connect to output pipe of chooser: %e", errPrefix, err)
		}

		err = cmd.Start()
		if err != nil {
			log.Fatalf("%sUnable to start chooser: %e", errPrefix, err)
		}
		pipe.Write([]byte(strings.Join(dirList, "\n")))
		pipe.Close()

		var pickedFile []byte
		pickedFile, err = ioutil.ReadAll(outPipe)
		if err != nil {
			log.Fatalf("%sUnable to read chooser output: %e", errPrefix, err)
		}
		err = cmd.Wait()
		if err != nil {
			log.Fatalf("%sCouldn't complete run of chooser: %e", errPrefix, err)
		}
		currentDirParts[len(currentDirParts)-1] = string(pickedFile)
	}
	papers, err := os.ReadDir(string(os.PathSeparator) + filepath.Join(currentDirParts...))
	if err != nil {
		log.Fatalf("%sUnable to get list of individual walpapers: %e", errPrefix, err)
	}
	var fullPath []string
	if *randomize {
		index := rand.Int() % len(papers)
		fullPath = append(currentDirParts, papers[index].Name())
	} else {
		//todo
	}

	var changeWalpaperFunc func(string) error
	switch runtime.GOOS {
	case "darwin":
		changeWalpaperFunc = func(s string) error {
			cmd := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"Finder\" to set desktop picture to POSIX file \"%s\"", s))
			return cmd.Run()

		}
	default:
		changeWalpaperFunc = func(string) error {
			return fmt.Errorf("buy a real computer")
		}
	}
	newWalpaper := string(os.PathSeparator) + filepath.Join(fullPath...)
	err = changeWalpaperFunc(newWalpaper)
	if err != nil {
		log.Fatalf("%sunable to change walpaper: %e", errPrefix, err)
	}
	f, err := os.Create(*stateFile)
	if err != nil {
		log.Fatalf("%sCreation of state file failed: %e", errPrefix, err)
	}
	f.Write([]byte(newWalpaper))
}
