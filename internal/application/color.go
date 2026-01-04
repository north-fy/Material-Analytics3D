package application

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/north-fy/Material-Analytics3D/internal/render"
)

func (r *Router) createColorScreen() fyne.CanvasObject {
	label := widget.NewLabel("Выберите цвет")

	colors := render.Colors
	var namesColor []string

	for k, _ := range colors {
		namesColor = append(namesColor, k)
	}

	colorSelect := widget.NewSelect(namesColor, func(name string) {
		r.managerScreen.settings["color"] = name
		r.managerScreen.setCurrentScreen("base")
	})

	cont := container.NewVBox(
		label,
		colorSelect,
	)

	return container.NewCenter(cont)
}
