package app

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/north-fy/Material-Analytics3D/internal/app/start"
)

func App() {
	log.Println("welcome!")
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())

	w := start.NewMainWindow(a)
	w.SetupUI(start.AuthObjects()...)
	w.Window.ShowAndRun()
}
