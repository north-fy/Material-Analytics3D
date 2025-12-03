package application

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (r *Router) createMainScreen() fyne.CanvasObject {

	// меню логин пароль + кнопки авторизации/регистрации

	entryLogin := widget.NewEntry()
	entryLogin.SetPlaceHolder("Логин")

	entryPassword := widget.NewPasswordEntry()
	entryPassword.SetPlaceHolder("Пароль")

	buttonAuth := widget.NewButton("Войти в аккаунт", func() {
		r.handleAuth(entryLogin.Text, entryPassword.Text)
	})

	buttonReg := widget.NewButton("Зарегистрироваться", func() {
		r.managerScreen.setCurrentScreen("reg")
	})

	contAdaptiveMenu := container.NewGridWrap(fyne.NewSize(700/2, 40), // FIX size
		entryLogin,
		entryPassword,
		buttonAuth,
		buttonReg,
	)

	contAdaptiveMenu = container.NewCenter(contAdaptiveMenu)

	// лого

	image, err := fyne.LoadResourceFromPath("./assets/Gopher.png")
	if err != nil {
		log.Fatal(err)
	}

	icon := widget.NewIcon(image)
	contLogo := container.NewGridWrap(fyne.NewSize(700/2, 180), // FIX size
		icon,
	)

	contLogo = container.NewCenter(contLogo)

	return container.NewVBox(
		contLogo,
		contAdaptiveMenu,
	)
}
