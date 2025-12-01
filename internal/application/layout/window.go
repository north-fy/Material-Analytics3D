package layout

import (
	"errors"
	"log"
)

func NewSpecificWindow(windows ...*Window) *SpecificWindow {
	ws := make(map[string]*Window, len(windows))
	for _, v := range windows {
		ws[v.Name] = v
	}

	result := &SpecificWindow{
		SpecificWindows: ws,
	}

	return result
}

func (sw *SpecificWindow) ViewObj(name string) error {
	window, ok := sw.SpecificWindows[name]
	if !ok {
		return errors.New("name is not available in map")
	}

	for _, v := range window.Objects {
		if v.Visible() == true {
			v.Show()
		} else {
			v.Hide()
		}

		log.Println("ViewObj", v.Container.Visible())
	}

	return nil
}
