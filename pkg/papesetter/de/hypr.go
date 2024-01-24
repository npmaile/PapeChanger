package de

import (
  "os/exec"
)

type Hypr struct{}

func(Hypr) SetPape( s string)error{
cmd := exec.Command("swww", "img", s)
return cmd.Run()
}
