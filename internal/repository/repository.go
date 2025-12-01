package repository

import (
	"database/sql"

	"github.com/north-fy/Material-Analytics3D/internal/repository/user"
)

func (d *Database) AddUser(u user.User) error {
	db := d.DB
	_, err := db.Exec(`INSERT OR IGNORE INTO users (login, password, access) VALUES ($1, $2, $3)`, u.Login, u.Password, u.Access)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) IsUser(u user.User) error {
	db := d.DB
	var login, password string
	_ = db.QueryRow(`SELECT login, password FROM users WHERE login=$1, password=$2`, u.Login, u.Password).Scan(&login, &password)
	if login == "" && password == "" {
		return sql.ErrNoRows
	}

	return nil
}

func (d *Database) GetUser(login, password string) (user.User, error) {
	unicUser := user.User{}
	db := d.DB
	row := db.QueryRow(`SELECT login, access FROM users WHERE login=$1, password=$2`, login, password).Scan(&unicUser.Login, &unicUser.Access)
	if row == nil {
		return unicUser, sql.ErrNoRows
	}

	return unicUser, nil
}

func (d *Database) UpdateUser(u user.User) error {
	db := d.DB
	_, err := db.Exec(`UPDATE users
	SET password = $1, access = $2
	WHERE login = $3`, u.Password, u.Access, u.Login)
	if err != nil {
		return err
	}

	return nil
}
