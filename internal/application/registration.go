package application

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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
			if ok := r.repo.IsUser(entryLogin.Text); ok {
				dialog.ShowError(errLoginIsExists, r.managerScreen.window)
				return
			}

			err := r.handleReg(entryLogin.Text, entryPassword.Text)
			if err != nil {
				dialog.ShowError(err, r.managerScreen.window)
				return
			}
		} else {
			dialog.ShowError(errNotCorrectPasswords, r.managerScreen.window)
			return
		}
	})

	buttonBack := widget.NewButton("В главное меню", func() {
		r.managerScreen.setCurrentScreen("auth")
	})

	contAdaptiveMenu := container.NewGridWrap(fyne.NewSize(700/2, 40), // FIX size
		entryLogin,
		entryPassword,
		entryPasswordTwo,
		buttonReg,
		buttonBack,
	)

	contAdaptiveMenu = container.NewCenter(contAdaptiveMenu)

	// лого

	image, err := fyne.LoadResourceFromPath("./assets/logo.png")
	if err != nil {
		log.Fatal(err)
	}

	icon := widget.NewIcon(image)
	contLogo := container.NewGridWrap(fyne.NewSize(700/2, 180), // FIX size
		icon,
	)

	labelReg := widget.NewLabel("Registration")
	labelReg.Resize(fyne.NewSize(200, 200))

	contLabelReg := container.NewCenter(labelReg)

	cont := container.NewCenter(container.NewVBox(
		contLogo,
		contLabelReg,
		contAdaptiveMenu,
	))

	return cont
}
