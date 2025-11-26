package start

import (
	"fyne.io/fyne/v2"
	app2 "github.com/north-fy/Material-Analytics3D/internal/app"
)

type MainWindow struct {
	Window fyne.Window
	App    fyne.App
}

func NewMainWindow(app fyne.App) *MainWindow {
	window := app.NewWindow("MaterialAnalytics3D - start")
	window.Resize(fyne.Size{Height: app2.AppHeight, Width: 700})
	window.SetFixedSize(true)

	return &MainWindow{
		Window: window,
		App:    app,
	}
}

func (mw *MainWindow) SetupUI(objects ...fyne.CanvasObject) {
	for _, obj := range objects {
		mw.Window.SetContent(obj)
	}
}
