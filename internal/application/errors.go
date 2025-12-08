package application

import "errors"

var (
	errWrongData           = errors.New("Неверный логин или пароль!")
	errNotCorrectPasswords = errors.New("Пароли не совпадают!")
	errLoginIsExists       = errors.New("Данный логин уже занят!")
)
