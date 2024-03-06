package main

import (
	"flag"
	"fmt"
	"os"
)

// [none] - change wallpaper
// --order=[random(default)|reverse|in-order]
// --cmd=[$command you want to run to select the wallpaper]
// -restore
// (positional)(optional) path to a specific wallpaper

// daemon - run daemon
// --order=[random(default)|reverse|in-order]
// --cmd=[$command you want to run to select the wallpaper]
// --selector=[we will pipe a list of directories (one per line) into this command, and change papechanger to the specific one]

// get - get currently set wallpaper
// --dir - get the directory of the wallpaper
// --dirs - get a listing of directories to pull wallpapers from

// cd - change directory wallpapers are being pulled from
// --no-change-pape - selectiion will take effect
// --dir - set the directory directly and skip the selection mechanism
// --cmd=[$command you want to run to select the wallpaper]
// --order=[random(default)|reverse|in-order]
// --selector=[we will pipe a list of directories (one per line) into this command, and change papechanger to the specific one]

func main() {
	if len(os.Args) < 2 {
		base()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "daemon":
		daemon()
	case "get":
		get()
	case "restore":
		restore()
	case "cd":
		cd()
	default:
		base()
	}

}

func base() {
	flag := flag.NewFlagSet("PapeChanger", flag.CommandLine.ErrorHandling())
	order := flag.String("order", "random", "order of papechanger to traverse the directory of wallpapers selected [random (default)|ordered]")
	restore := flag.Bool("restore", false, "restore last set wallpaper instead of finding another")
	cmd := flag.String("cmd", "", "command to be run to set the wallpaper. This string is passed directly to the OS default shell with any instance of %s replaced with the name of the wallpaper")
	err := flag.Parse(os.Args[1:])
	if err != nil {
		fmt.Println("todo")
		os.Exit(1)
	}

}
func daemon() {
	flag := flag.NewFlagSet("daemon", flag.CommandLine.ErrorHandling())
	order := flag.String("order", "random", "order of papechanger to traverse the directory of wallpapers selected [random (default)|ordered]")
	cmd := flag.String("cmd", "", "command to be run to set the wallpaper. This string is passed directly to the OS default shell with any instance of %s replaced with the name of the wallpaper")
	selectorcmd := flag.String("selector", "", "Command to be run to select the directory to pull wallpapers from. A list of directories will be passed to it on stdin, and whatever directory comes back from it will be selected by papeChanger")
	err := flag.Parse(os.Args[2:])
	if err != nil {
		fmt.Println("todo")
		os.Exit(1)
	}
}
func get() {
	flag := flag.NewFlagSet("get", flag.CommandLine.ErrorHandling())
	dirs := flag.Bool("dirs", false, "get the directory that is being used to find directories")
	dir := flag.Bool("dir", false, "get the directory that wallpapers are being pulled from right now")
	err := flag.Parse(os.Args[2:])
	if err != nil {
		fmt.Println("todo")
		os.Exit(1)
	}

}
func cd() {
	flag := flag.NewFlagSet("cd", flag.CommandLine.ErrorHandling())
	order := flag.String("order", "random", "order of papechanger to traverse the directory of wallpapers selected [random (default)|ordered]")
	selectorcmd := flag.String("selector", "", "Command to be run to select the directory to pull wallpapers from. A list of directories will be passed to it on stdin, and whatever directory comes back from it will be selected by papeChanger")
	cmd := flag.String("cmd", "", "command to be run to set the wallpaper. This string is passed directly to the OS default shell with any instance of %s replaced with the name of the wallpaper")
	selectorcmd := flag.String("selector", "", "Command to be run to select the directory to pull wallpapers from. A list of directories will be passed to it on stdin, and whatever directory comes back from it will be selected by papeChanger")

	err := flag.Parse(os.Args[2:])
	if err != nil {
		fmt.Println("todo")
		os.Exit(1)
	}
}
