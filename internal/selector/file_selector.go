package selector

import (
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"strings"
	"time"
)

var gRand *rand.Rand

func init() {
	gRand = rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
}

func SelectWallpaperRandom(papeDir string) (string, error) {
	papeCandidates, err := os.ReadDir(papeDir)
	if err != nil {
		return "", err
	}
	var papesToChooseFrom []fs.DirEntry
	for _, f := range papeCandidates {
		if f.Type().IsRegular() {
			papesToChooseFrom = append(papesToChooseFrom, f)
		}
	}
	index := gRand.Int() % len(papesToChooseFrom)

	ret := fmt.Sprintf("%s%s%s", papeDir, string(os.PathSeparator), papesToChooseFrom[index].Name())
	return ret, nil
}

func SelectWallpaperInOrder(papeDir string, currentWallpaperFullPath string) (string, error) {
	papeCandidates, err := os.ReadDir(papeDir)
	if err != nil {
		return "", err
	}
	var papesToChooseFrom []fs.DirEntry
	for _, f := range papeCandidates {
		if f.Type().IsRegular() {
			papesToChooseFrom = append(papesToChooseFrom, f)
		}
	}

	var ret string

	fullCurPapePath := strings.Split(currentWallpaperFullPath, string(os.PathSeparator))
	currentPape := fullCurPapePath[len(fullCurPapePath)-1]

	for index, entry := range papesToChooseFrom {
		if entry.Name() == currentPape {
			if index >= len(papesToChooseFrom)-1 {
				ret = papesToChooseFrom[0].Name()
			} else {
				ret = papesToChooseFrom[index+1].Name()
			}
		}
	}
	if ret == "" {
		//current wallpaper was not found, therefore just use the first one
		ret = papesToChooseFrom[0].Name()
	}
	realret := papeDir + string(os.PathSeparator) + ret
	return realret, nil
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
	return dirs, nil
}
