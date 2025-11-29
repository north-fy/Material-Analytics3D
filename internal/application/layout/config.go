package layout

import "fyne.io/fyne/v2"

const (
	AppHeight float32 = 500.0
	AppWidth  float32 = 700.0
)

//type specificWindow interface {
//	CreateContainer()
//}

type Object struct {
	Cont *fyne.Container
	Data map[string]interface{}
}

type Window struct {
	Name    string
	Objects []Object
	IsShow  bool
}

type SpecificWindow struct {
	specificWindows map[string]*Window
}
