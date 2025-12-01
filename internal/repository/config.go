package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SSLmode  string `yaml:"sslmode"`
}

type Database struct {
	DB *sql.DB
}

//func NewConfig() (*Config, error) {
//	data, err := ioutil.ReadFile("config.yaml")
//	if err != nil {
//		return nil, err
//	}
//
//	var config Config
//	err = yaml.Unmarshal(data, &config)
//	if err != nil {
//		return nil, err
//	}
//
//	return &config, nil
//}

func InitDB(cfg Config) (*Database, error) {
	dataSourceName := fmt.Sprintf("host=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.SSLmode)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users(
    id SERIAL,
    login TEXT,
    password TEXT,
    access INT)`)
	if err != nil {
		return nil, err
	}

	return &Database{DB: db}, nil
}
