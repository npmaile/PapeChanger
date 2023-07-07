package de

import (
	"fmt"
	"os/exec"
)

type Plasma struct{}

func (Plasma) SetPape(s string) error {
	cmd := exec.Command("dbus-send", "--session", "--dest=org.kde.plasmashell", "--type=method_call", "/PlasmaShell", "org.kde.PlasmaShell.evaluateScript", fmt.Sprintf(`string:
var Desktops = desktops();
for (i=0;i<Desktops.length;i++) {
        d = Desktops[i];
        d.wallpaperPlugin = "org.kde.image";
        d.currentConfigGroup = Array("Wallpaper", "org.kde.image", "General");
        d.writeConfig("Image", "file://%s");
}`, s))
	return cmd.Run()

}
