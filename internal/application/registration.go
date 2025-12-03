package application

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (r *Router) createRegScreen() fyne.CanvasObject {

	// меню логин пароль + кнопки авторизации/регистрации

	entryLogin := widget.NewEntry()
	entryLogin.SetPlaceHolder("Логин")

	entryPassword := widget.NewPasswordEntry()
	entryPassword.SetPlaceHolder("Пароль")

	entryPasswordTwo := widget.NewPasswordEntry()
	entryPasswordTwo.SetPlaceHolder("Подтвердите пароль")

	buttonReg := widget.NewButton("Зарегистрироваться", func() {
		if entryPassword.Text == entryPasswordTwo.Text {
			if ok := r.repo.IsUser(entryLogin.Text); !ok {
				// сюда виджет
				log.Println("данный логин уже занят!")
			}

			r.handleReg(entryLogin.Text, entryPassword.Text)
		}
	})

	contAdaptiveMenu := container.NewGridWrap(fyne.NewSize(700/2, 40), // FIX size
		entryLogin,
		entryPassword,
		entryPasswordTwo,
		buttonReg,
	)

	contAdaptiveMenu = container.NewCenter(contAdaptiveMenu)

	// лого

	labelReg := widget.NewLabel("Registration")
	contLabelReg := container.NewGridWrap(fyne.NewSize(700/2, 180), // FIX size
		labelReg,
	)

	contLabelReg = container.NewCenter(contLabelReg)

	cont := container.NewCenter(container.NewVBox(
		contLabelReg,
		contAdaptiveMenu,
	))

	return cont
}
