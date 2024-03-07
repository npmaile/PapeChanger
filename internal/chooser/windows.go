//go:build windows
// +build windows

package chooser
func Chooser(dirs []string) (string, error) {
	return BuiltIn(dirs)
}
