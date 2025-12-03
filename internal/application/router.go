package application

import (
	"github.com/north-fy/Material-Analytics3D/internal/calculator"
	"github.com/north-fy/Material-Analytics3D/internal/repository"
)

func newRouter(cfg repository.Config) (*Router, error) {
	manager := windowManager()

	calcService := calculator.CreateCalcService()
	DB, err := repository.InitDB(cfg)
	if err != nil {
		return nil, err
	}

	return &Router{
		managerWindow: manager,
		calcService:   calcService,
		repo:          DB,
	}, nil
}

func (r *Router) route() error {
	for key, window := range r.managerWindow.SpecificWindows {
		switch key {
		case "auth":
			if window.Data["reg"] == true {
				login := window.Data["login"]
				password := window.Data["password"]

				// потом заменить на хендлер функцию выскакивания неправильного логина/пароля
				_, err := r.repo.GetUser(login.(string), password.(string))
				if err != nil {
					return err
				}
				// логика захода в base
				// ...
			}
		}
	}
	return nil
}
