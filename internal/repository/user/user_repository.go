package user

func NewUser(access AccessType, login, password string) (*User, error) {
	user := &User{
		Login:    login,
		Password: password,
		Access:   access,
	}

	return user, nil
}

func (u User) AuthUser(login, password string) error {
	if u.Login != login {
		return errWrongData
	}

	if u.Password != password {
		return errWrongData
	}

	return nil
}

func (u *User) UpdateAccessUser(access AccessType) {
	u.Access = access
}

func (u *User) CheckAccessUser(access AccessType) bool {
	if u.Access == access {
		return true
	}

	return false
}
