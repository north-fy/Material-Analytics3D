package main

import (
	"fyne.io/fyne/v2"
)

type MainWindow struct {
	window fyne.Window
	app    fyne.App
}

func NewMainWindow(app fyne.App) *MainWindow {
	window := app.NewWindow("MaterialAnalytics3D - main")
	window.Resize(fyne.Size{Height: 500, Width: 500})
	window.SetFixedSize(true)

	return &MainWindow{
		window: window,
		app:    app,
	}
}

func (mw *MainWindow) SetupUI(objects ...fyne.CanvasObject) {
	for _, obj := range objects {
		mw.window.SetContent(obj)
	}
}
