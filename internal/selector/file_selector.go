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
	return fmt.Sprintf("%s%s%s", papeDir, string(os.PathSeparator), papeCandidates[index].Name()), nil
}

func ListDirectories(dirOfDirs string) ([]string, error) {
	DirCandidates, err := os.ReadDir(dirOfDirs)
	if err != nil {
		return []string{""}, err
	}
	var dirs = make([]string, len(DirCandidates))
	for _, possibleDir := range DirCandidates {
		if possibleDir.IsDir() {
			dirs = append(dirs, possibleDir.Name())
		}
	}
	return dirs, nil
}
