//go:build windows
// +build windows

package chooser

func Chooser([]string) (string, error) {
	return "", fmt.Errorf("Jeez, I'm working on it!")
}
