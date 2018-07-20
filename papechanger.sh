#!/bin/bash


#declare the root directory for the pape folders
walpaperdir=""

#the script starts here
folderpath="$walpaperdir/$(cat $HOME/.papefolder)"

pickpape()
{
	selectionfile="$(ls "$folderpath" | shuf -n 1 )"
	selectionfullpath="$folderpath/$selectionfile"
	echo "$selectionfullpath"
	echo $selectionfullpath >> ~/.papehistory
}

changepape()
{
	numscreens="$(xrandr | grep " connected" | awk '{print $1}' | wc -l)"
	fehargs=("--bg-fill")
	while [ $numscreens -gt 0 ]
	do
		newarg="$(pickpape)"
		fehargs+='" "'
		fehargs+="$newarg"
		numscreens=$(($numscreens-1))
	done
	echo "fehargs= ""${fehargs[@]}"
	eval feh '"'$fehargs'"'
}

change_pape_folder()
{
	options=$(ls -d "$walpaperdir"/*/ | sed "s:\(\/home\/nate\/Pictures\/wallpapers\/\)\(.*\)\/:\2:")
	selection=$(echo "$options" | rofi -dmenu)
	echo $selection > ~/.papefolder
	folderpath="$walpaperdir/$(cat ~/.papefolder)"
	changepape
}

###############################
#main body
###############################

if [ -z "$*" ]; then
changepape
else
change_pape_folder
fi
