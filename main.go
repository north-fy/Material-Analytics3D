package main

import (
	"log"
	"os"

	"github.com/north-fy/Material-Analytics3D/internal/application"
	"github.com/north-fy/Material-Analytics3D/internal/repository"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Sslmode  string `yaml:"sslmode"`
	} `yaml:"database"`
	App struct {
		AppHeight float32 `yaml:"app-height"`
		AppWidth  float32 `yaml:"app-width"`
		Name      string  `yaml:"name"`
		FixedSize bool    `yaml:"fixed-size"`
	} `yaml:"app"`
}

func main() {
	cfg, err := LoadConfig("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	cfgDB := repository.NewConfig(cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.Sslmode)
	cfgApp := application.NewConfig(cfg.App.AppHeight, cfg.App.AppWidth, cfg.App.Name, cfg.App.FixedSize)

	app, err := application.NewMainApp(*cfgDB, *cfgApp)
	if err != nil {
		log.Fatal(err)
	}

	app.Run()
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var (
		config Config
	)

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, err
}
