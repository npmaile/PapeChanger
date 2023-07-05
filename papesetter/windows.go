//go:build windows
// +build windows

package papesetter

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32DLL           = windows.NewLazyDLL("user32.dll")
	procSystemParamInfo = user32DLL.NewProc("SystemParametersInfoW")
)

func SetPape(s string) error {
	path, err := windows.UTF16PtrFromString(s)
	if err != nil {
		return err
	}
	procSystemParamInfo.Call(20, 0, uintptr(unsafe.Pointer(path)), 0x001A)
	return nil
}
