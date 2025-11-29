package application

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/north-fy/Material-Analytics3D/internal/application/layout"
	"github.com/north-fy/Material-Analytics3D/internal/application/layout/authentication"
)

// что у нас по архитектуре?
// router - выбирает какое окно показывать
// manager - показывает/скрывает окна
//

type MainApp struct {
	// окна и сервисы для них, как реализовать?
	// интерфейс?
	fyneApp fyne.App
	router  *Router
	manager *Manager
}

func NewMainApp() *MainApp {
	a := app.New()
	return &MainApp{
		fyneApp: a,
	}

}

func App() {
	log.Println("welcome!")
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())

	w := layout.NewMainWindow(a)
	w.SetupUI(authentication.AuthObjects())
	w.Window.ShowAndRun()
}
