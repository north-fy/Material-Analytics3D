package layout

import (
	"errors"

	"fyne.io/fyne/v2"
)

func NewSpecificWindow(App fyne.App, windows ...Window) *SpecificWindow {
	window := App.NewWindow("Material-Analytics3D")
	window.Resize(fyne.Size{Height: AppHeight, Width: AppWidth})
	window.SetFixedSize(true)

	ws := make(map[string]*Window, len(windows))
	for _, v := range windows {
		ws[v.Name] = &v
	}

	result := &SpecificWindow{
		specificWindows: ws,
	}

	return result
}

func (sw *SpecificWindow) ViewObj(name string) error {
	if _, ok := sw.specificWindows[name]; !ok {
		return errors.New("name is not available in map")
	}

	for _, v := range sw.specificWindows[name].Objects {
		if v.Cont.Visible() == false {
			v.Cont.Show()
		} else {
			v.Cont.Hide()
		}
	}

	return nil
}
