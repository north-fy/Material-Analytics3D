package application

import (
	"github.com/north-fy/Material-Analytics3D/internal/application/layout"
	"github.com/north-fy/Material-Analytics3D/internal/calculator"
	"github.com/north-fy/Material-Analytics3D/internal/repository"
)

type ConfigApp struct {
	AppHeight float32 `yaml:"app-height"`
	AppWidth  float32 `yaml:"app-width"`
	Name      string  `yaml:"name"`
	FixedSize bool    `yaml:"fixed-size"`
}

type Router struct {
	managerWindow *layout.SpecificWindow
	calcService   *calculator.CalcService
	repo          *repository.Database
}

// managerService map[string]interface{}
