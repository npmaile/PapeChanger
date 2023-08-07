
package de

import (
	"os/exec"
  "fmt"
  "os"
  "strings"
)

type I3 struct{}

func (I3) SetPape(s string) error {
  pathVar, present := os.LookupEnv("PATH")
  if !present {
    return fmt.Errorf("PATH not found")
  }
  paths := strings.Split(pathVar, ":")


  for _, path := range paths {
    files, err := os.ReadDir(path)
    if err != nil {
      continue
    }
    for _, file := range files {
      switch file.Name() {
      case "feh":
        cmd := exec.Command("feh", "--bg-fill", s)
        return cmd.Run()
      }
    }
  }
  return fmt.Errorf("Unable to determine wallpaper setter i3")
}
