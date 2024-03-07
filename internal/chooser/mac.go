//go:build darwin
// +build darwin

package chooser

import "fmt"

func Chooser(directories []string) (string, error) {
	fmt.Println("returning built-in chooser")
	return BuiltIn(directories)
}
