package repository

import (
	"database/sql"

	"github.com/north-fy/Material-Analytics3D/internal/repository/user"
)

func (d *Database) AddUser(u user.User) error {
	db := d.DB
	_, err := db.Exec(`INSERT INTO users (login, password, access) VALUES ($1, $2, $3)`, u.Login, u.Password, u.Access.Access) // FIX
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) IsUser(login string) bool {
	db := d.DB
	u := user.User{}
	id := 0
	_ = db.QueryRow(`SELECT * FROM users WHERE login=$1`, login).Scan(&id, &u.Login, &u.Password, &u.Access.Access)

	if u.Login == "" || login == "" {
		return false
	}

	return true
}

func (d *Database) GetUser(login string) (user.User, error) {
	db := d.DB
	unicUser := user.User{}
	row := db.QueryRow(`SELECT login, password, access FROM users WHERE login=$1`, login).Scan(&unicUser.Login, &unicUser.Password, &unicUser.Access.Access)
	if row != nil {
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
