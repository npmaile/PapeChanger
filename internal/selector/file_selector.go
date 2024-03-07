package selector

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

var gRand *rand.Rand

func init() {
	gRand = rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
}

func SelectWallpaperRandom(papeDir string) (string, error) {
	papeCandidates, err := ListFiles(papeDir)
	if err != nil {
		return "", err
	}
	index := gRand.Int() % len(papeCandidates)

	ret := fmt.Sprintf("%s%s%s", papeDir, string(os.PathSeparator), papeCandidates[index])
	fmt.Printf("something something words: %s", ret)
	return ret, nil
}

func SelectWallpaperInOrder(papeDir string, currentWallpaperFullPath string) (string, error) {
	papeCandidates, err := ListFiles(papeDir)
	if err != nil {
		return "", err
	}
	var ret string

	fullCurPapePath := strings.Split(currentWallpaperFullPath, string(os.PathSeparator))
	currentPape := fullCurPapePath[len(fullCurPapePath)-1]

	for index, entry := range papeCandidates {
		if entry == currentPape {
			if index >= len(papeCandidates)-1 {
				ret = papeCandidates[0]
			} else {
				ret = papeCandidates[index+1]
			}
		}
	}
	if ret == "" {
		//current wallpaper was not found, therefore just use the first one
		ret = papeCandidates[0]
	}
	realret := papeDir + string(os.PathSeparator) + ret
	return realret, nil
}

func ListFiles(directory string) ([]string, error) {
	fileCandidates, err := os.ReadDir(directory)
	if err != nil {
		return []string{""}, err
	}
	var files = make([]string, 0)
	for _, possibleFile := range fileCandidates {
		if !possibleFile.IsDir() {
			files = append(files, possibleFile.Name())
		}
	}
	return files, nil
}
func ListDirectories(dirOfDirs string) ([]string, error) {
	DirCandidates, err := os.ReadDir(dirOfDirs)
	if err != nil {
		return []string{""}, err
	}
	var dirs = make([]string, 0)
	for _, possibleDir := range DirCandidates {
		if possibleDir.IsDir() {
			dirs = append(dirs, possibleDir.Name())
		}
	}
	collate.New(language.English).SortStrings(dirs)
	return dirs, nil
}
