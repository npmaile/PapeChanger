//go:build windows
// +build windows

package chooser

func Chooser(dirs []string) (string, error) {
	// will replace with something else if I find a good chooser for windows
	return BuiltIn(dirs)
}
