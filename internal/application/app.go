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

	window := App.NewWindow(cfgApp.Name)
	window.Resize(fyne.Size{Height: cfgApp.AppHeight, Width: cfgApp.AppWidth})
	window.SetFixedSize(cfgApp.FixedSize)

	router, err := newRouter(cfgRepo)
	if err != nil {
		return nil, err
	}

	window.SetContent(router.managerWindow.SpecificWindows["auth"].Objects[0].Container)

	err = router.managerWindow.ViewObj("auth")
	if err != nil {
		return nil, err
	}

	window.Show()

	return &MainApp{
		FyneApp: App,
		Router:  router,
	}, nil
}

func (mp *MainApp) Run() error {
	err := mp.Router.route()
	if err != nil {
		return err
	}

	mp.FyneApp.Run()

	return nil
}
