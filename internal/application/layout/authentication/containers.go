package authentication

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/north-fy/Material-Analytics3D/internal/application/layout"
)

func (w Window) NewAuthMenu() {
	data := make(map[string]interface{})

	entryLogin := widget.NewEntry()
	entryLogin.SetPlaceHolder("Логин")

	entryPassword := widget.NewPasswordEntry()
	entryPassword.SetPlaceHolder("Пароль")

	buttonAuth := widget.NewButton("Войти в аккаунт", func() {
		data["login"] = entryLogin.Text
		data["password"] = entryPassword.Text
	})

	buttonReg := widget.NewButton("Зарегистрироваться", func() {
		data["reg"] = true
	})

	cont := container.NewGridWrap(fyne.Size{Width: layout.AppWidth / 2, Height: 40},
		entryLogin,
		entryPassword,
		buttonAuth,
		buttonReg,
	)
	cont = container.NewCenter(cont)
	obj := Object{
		Cont: cont,
		Data: data,
	}

	w.Objects = append(w.Objects, obj)
}

func (w *Window) NewLogo() {
	image, err := fyne.LoadResourceFromPath("./assets/Gopher.png")
	if err != nil {
		log.Fatal(err)
	}

	icon := widget.NewIcon(image)
	cont := container.NewGridWrap(fyne.Size{Width: layout.AppWidth / 1.5, Height: 180},
		icon,
	)

	obj := Object{
		Cont: cont,
		Data: nil,
	}

	w.Objects = append(w.Objects, obj)
}
