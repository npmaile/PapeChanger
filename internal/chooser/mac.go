//go:build darwin
// +build darwin

package chooser

func Chooser(directories []string) (string, error) {
	return BuiltIn(directories)
}
