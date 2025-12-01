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
	*fyne.Container
}

type Window struct {
	Name    string
	Objects []*Object
	Data    map[string]interface{}
	IsShow  bool
}

type SpecificWindow struct {
	SpecificWindows map[string]*Window
}
