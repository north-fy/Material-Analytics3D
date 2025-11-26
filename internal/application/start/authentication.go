package start

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func AuthObjects() []fyne.CanvasObject {
	entryLogin := widget.NewEntry()
	entryLogin.Resize(fyne.Size{Width: 200, Height: 100})
	entryLogin.Move(fyne.Position{X: })

	objects := []fyne.CanvasObject{
		entryLogin,
		widget.NewPasswordEntry(),
	}

	return objects
}
