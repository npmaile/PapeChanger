package selector

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func SelectWallpaper(papeDir string) (string, error) {
	papeCandidates, err := os.ReadDir(papeDir)
	if err != nil {
		return "", err
	}
	index := rand.Int() % len(papeCandidates)
	return fmt.Sprintf("%s%s%s", papeDir, os.PathSeparator, papeCandidates[index].Name), nil
}

func SelectDirectory(dirOfDirs string, selectionFunc func([]string) (string, error)) (string, error) {
	DirCandidates, err := os.ReadDir(dirOfDirs)
	if err != nil {
		return "", err
	}
	var dirs = make([]string, len(DirCandidates))
	for _, possibleDir := range DirCandidates {
		if possibleDir.IsDir() {
			dirs = append(dirs, possibleDir.Name())
		}
	}
	selection, err := selectionFunc(dirs)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s%s", dirOfDirs, os.PathSeparator, selection), nil

}
