# PapeChanger

This program switches your wallpaper to a random selection from a directory and allows changing the directory. It's meant to be invoked via a keyboard macro, though it will work from the command line.

<https://github.com/npmaile/PapeChanger/assets/17767713/1cc584cc-681c-4020-914f-be48e6bd9b4d>

# How to use

This program assumes you have a directory filled with directories filled with walpapers.
If your walpapers are laid out in another configuration, You won't be able to use the directory switching capabilities.
If your walpapers are laid out another way, you may wish to symlink your wallpapers into a directory structure that is compliant or look elsewhere for walpaper management

Once built, The program should be placed somehwere in your `$PATH` so it can be called.
Set it up by running `papeChanger -setup $SOME_PATH_TO_SOME_WALLPAPER`, which should set your walpaper
Test the directory changing functionality by running `papeChanger -c`. If it does not bring up a window to select your walpaper directory, try running it as `papechanger -c --useBuiltin` to use the built in directory switcher.
Once you've confirmed it works on your system, you will want to bind `papeChanger` to a hotkey and `papeChanger -c` to another hotkey so you can change your walpapers without needing to open a terminal or run dialogue.

# Installation
## MacOS/Windows
If you are on MacOS or Windows, you can install papeChanger by downloading an installable build on the [releases](https://github.com/npmaile/PapeChanger/releases) page. This will give you the PapeChanger Desktop Application which is installed and run in the way you would expect.

## Command line
On all supported platforms, the core functionality of PapeChanger is available by running `go install github.com/npmaile/papeChanger@latest` in the terminal (a functioning Go environment is necessary) Once enough people complain, I'll probably make builds for all the systems available on the releases page like the packaged ones.

# Building
## Minimum Requirements

- A functioning go environment
- All of the [requirements](https://developer.fyne.io/started/) to build the [fyne](https://fyne.io/) toolkit

## Minimum Requirements (lite edition)
- A functioning go environment
- Rofi installed
can be built by running `cd extra/papeChanger-lite && CGO_ENABLED=0 go build -o papechanger`

## Windows

It has a graphical component that requires a c compiler on windows
Once you have one installed (only on windows), you can simply build using `go build -o papeChanger main.go`

# Supported Systems

I will be supporting systems as I become aware of them. The two things necessary for a supported system are a mechanism in code to change the wallpaper and a window to choose which directory your walpapers are in.

- \[x\] Windows
- \[x\] Mac
- \[x\] Linux/Sway
- \[x\] Linux/i3wm
- \[x\] Linux/XFCE
- \[x\] Linux/KDE
- \[x\] Linux/Gnome
- \[x\] Linux/Hyprland
- \[ \] Linux/x11-generic
- \[ \] Linux/Mate
- \[ \] Linux/Cinnamon
- \[ \] Linux/Budgie
- \[ \] Linux/LxQt
- \[ \] Linux/Deepin
