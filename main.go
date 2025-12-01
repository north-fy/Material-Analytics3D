package main

import (
	"log"
	"os"

	"github.com/north-fy/Material-Analytics3D/internal/application"
	"github.com/north-fy/Material-Analytics3D/internal/repository"
	"gopkg.in/yaml.v3"
)

type Config struct {
	dbCfb  repository.Config     `yaml:"database"`
	appCfg application.ConfigApp `yaml:"app"`
}

func main() {
	cfg, err := LoadConfig("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	app, err := application.NewMainApp(cfg.dbCfb, cfg.appCfg)
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	log.Println(data)
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
