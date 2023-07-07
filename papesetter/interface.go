package papesetter

import "errors"

type DE interface {
	SetPape(string) error
}

type deNotFound struct{}

func (deNotFound) SetPape(string) error {
	return errors.New("your operating system is not configured. If you can, please make an issue at https://github.com/npmaile/papeChanger")
}
