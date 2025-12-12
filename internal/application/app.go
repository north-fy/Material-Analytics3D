package application

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/north-fy/Material-Analytics3D/internal/repository"
)

// что у нас по архитектуре?
// router - выбирает какое окно показывать
// manager - показывает/скрывает окна
//

type MainApp struct {
	FyneApp fyne.App
	Router  *Router
}

func NewMainApp(cfgRepo repository.Config, cfgApp ConfigApp) (*MainApp, error) {
	App := app.New()
	App.Settings().SetTheme(theme.DarkTheme())

	image, err := fyne.LoadResourceFromPath("./assets/logo.png")
	if err != nil {
		return nil, err
	}

	App.SetIcon(image)

	window := App.NewWindow(cfgApp.Name)
	window.Resize(fyne.Size{Height: cfgApp.AppHeight, Width: cfgApp.AppWidth})
	window.SetFixedSize(cfgApp.FixedSize)

	router, err := newRouter(cfgRepo, window)
	if err != nil {
		return nil, err
	}

	mainApp := &MainApp{
		FyneApp: App,
		Router:  router,
	}

	mainApp.initScreens()

	return mainApp, nil
}

func (mp *MainApp) initScreens() {
	mp.Router.managerScreen.addScreen("auth", mp.Router.createMainScreen())
	mp.Router.managerScreen.addScreen("reg", mp.Router.createRegScreen())
	mp.Router.managerScreen.addScreen("base", mp.Router.createBaseScreen())

	mp.Router.managerScreen.setCurrentScreen("auth")
}

func (mp *MainApp) Run() {
	mp.Router.managerScreen.window.ShowAndRun()
}
