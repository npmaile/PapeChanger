# PapeChanger
Recently updated to lightnign fast go!
Nearly a 10% improvement in speed (on macos) (compared to shell version)

# How to use
This program assumes you have a directory filled with directories filled with walpapers. 
If your walpapers are laid out in another configuration, You won't be able to use the directory switching capabilities.
If your walpapers are laid out another way, you may wish to symlink your wallpapers into a directory structure that is compliant or look elsewhere for walpaper management

Once built, The program should be placed somehwere in your `$PATH` so it can be called. 
Set it up by running `papeChanger -setup $SOME_PATH_TO_SOME_WALLPAPER`, which should set your walpaper
Test the directory changing functionality by running `papeChanger -c`. If it does not bring up a window to select your walpaper directory, try running it as `papechanger -c --useBuiltin` to use the built in directory switcher. 
Once you've confirmed it works on your system, you will want to bind `papeChanger` to a hotkey and `papeChanger -c` to another hotkey so you can change your walpapers without needing to open a terminal or run dialogue. 

# Building
It has a graphical component that requires a c compiler on windows
Once you have one installed (only on windows), you can simply build using `go build -o papeChanger main.go`

# Supported Systems
I will be supporting systems as I become aware of them. The two things necessary for a supported system are a mechanism in code to change the wallpaper and a window to choose which directory your walpapers are in. 
- [x] Windows
- [x] Mac
- [x] Linux/Sway
- [ ] Linux/x11
- [ ] BSD/x11
