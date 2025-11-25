package repository

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	host     string `yaml:"host"`
	user     string `yaml:"user"`
	password string `yaml:"password"`
	sslmode  string `yaml:"sslmode"`
}
type Database struct {
	db *sql.DB
}

func NewConfig() (*Config, error) {
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func InitDB(cfg Config) (*Database, error) {
	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s sslmode=%s",
		cfg.host, cfg.user, cfg.password, cfg.sslmode)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	
	return &Database{db: db}, nil
}
