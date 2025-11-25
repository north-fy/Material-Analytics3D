package user

func NewUser(access AccessType, email, password string) (*User, error) {
	user := &User{
		Email:    email,
		Password: password,
		Access:   access,
	}

	return user, nil
}

func AuthUser() {

}

func (u *User) UpdateAccessUser(access AccessType) {
	u.Access = access
}
