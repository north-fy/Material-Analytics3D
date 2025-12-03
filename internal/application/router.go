package application

import (
	"log"

	"fyne.io/fyne/v2"
	"github.com/north-fy/Material-Analytics3D/internal/calculator"
	"github.com/north-fy/Material-Analytics3D/internal/repository"
)

type Router struct {
	managerScreen *ScreenManager
	calcService   *calculator.CalcService
	repo          *repository.Database
}

func newRouter(cfg repository.Config, window fyne.Window) (*Router, error) {
	manager := NewScreenManager(window)
	calcService := calculator.CreateCalcService()
	DB, err := repository.InitDB(cfg)
	if err != nil {
		return nil, err
	}

	return &Router{
		managerScreen: manager,
		calcService:   calcService,
		repo:          DB,
	}, nil
}

func (r *Router) handleAuth(login, password string) {
	u, err := r.repo.GetUser(login, password)
	if err != nil {
		// ЗАМЕНИТЬ
		log.Fatal(err)
		return
	}

	if u.Login == "" || u.Password == "" {
		// тут мб виджет
		log.Println("agaaga")
		return
	}
	_ = u

	log.Println("handled!")
	//switchTo(base, access)
}
