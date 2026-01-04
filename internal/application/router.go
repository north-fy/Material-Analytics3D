package application

import (
	"fyne.io/fyne/v2"
	"github.com/north-fy/Material-Analytics3D/internal/calculator"
	"github.com/north-fy/Material-Analytics3D/internal/repository"
	"github.com/north-fy/Material-Analytics3D/internal/repository/user"
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

func (r *Router) handleAuth(login, password string) error {
	u, err := r.repo.GetUser(login)
	if err != nil {
		return errWrongData
	}

	if u.Login == "" || u.Password == "" {
		return errWrongData
	}

	if login == u.Login && password == u.Password {
		r.managerScreen.setCurrentScreen("base")
		return nil
	}

	return errWrongData
}

func (r *Router) handleReg(login, password string) error {
	u, err := user.NewUser(user.AccessType{Access: user.AccessUser}, login, password)
	if err != nil {
		return err
	}

	err = r.repo.AddUser(*u)
	if err != nil {
		return err
	}

	return nil
}
