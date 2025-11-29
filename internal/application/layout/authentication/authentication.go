package authentication

import (
	"fyne.io/fyne/v2"
)

type Object struct {
	Cont *fyne.Container
	Data map[string]interface{}
}

type Window struct {
	Name    string
	Objects []Object
	IsShow  bool
}

func NewAuthWindow() *Window {
	window := &Window{Name: "auth"}
	window.NewLogo()
	window.NewAuthMenu()

	return window
}
