# PapeChanger
This is where I am keeping my walpaper changing scripts

depends on:
rofi (dmenu should work if you replace the rofi call with dmenu)
feh for setting the background.

to use it, you need to first go into the script and change the location of the wallpaper root directory. This is where the directories containing your images should be kept. 

Usage: if you run it with no arguments, it will change your walpaper to a random one from the directory you've chosen. you can choose a directory by running it with literally any argument. 

Features:
picking a different image for each monitor attached(kill me)
utilizing rofi for some dumb ass reason
literally nothing else. Please don't use this script
possible arbitrary code execution due to the eval statement someplace in there. 
logging your papes in ~/.papehistory so you can go back and look at ones that you used in the past (if you use feh $(tail ~/.papehistory | rofi -dmenu), you can pick a recently viewed pape and look at it.)

upcoming:
a pape downloader to grab walpapers from imageboards like 4chan or 8chan or other places I'm sure
more concise programming
support for other things I'm sure


moonshots:
support for moving walpapers in addition to just the static ones
support for various wayland desktops

