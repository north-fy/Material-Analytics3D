package layout

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewAuthWindow() *Window {
	window := &Window{
		Name:    "auth",
		Objects: []*Object{},
		Data:    make(map[string]interface{}),
		IsShow:  false,
	}

	window.newLogo()
	window.newAuthMenu()

	return window
}

// надо завтра рефакторить!

func (w *Window) newAuthMenu() {
	entryLogin := widget.NewEntry()
	entryLogin.SetPlaceHolder("Логин")

	entryPassword := widget.NewPasswordEntry()
	entryPassword.SetPlaceHolder("Пароль")

	buttonAuth := widget.NewButton("Войти в аккаунт", func() {
		login := entryLogin.Text
		password := entryPassword.Text
		//... auth()
	})

	buttonReg := widget.NewButton("Зарегистрироваться", func() {
		w.Data["reg"] = true
	})

	cont := container.NewGridWrap(fyne.Size{Width: AppWidth / 2, Height: 40},
		entryLogin,
		entryPassword,
		buttonAuth,
		buttonReg,
	)
	cont = container.NewCenter(cont)

	w.Objects = append(w.Objects, &Object{cont})
}

func (w *Window) newLogo() {
	image, err := fyne.LoadResourceFromPath("./assets/Gopher.png")
	if err != nil {
		log.Fatal(err)
	}

	icon := widget.NewIcon(image)
	cont := container.NewGridWrap(fyne.Size{Width: AppWidth / 1.5, Height: 180},
		icon,
	)

	cont = container.NewCenter(cont)

	w.Objects = append(w.Objects, &Object{cont})
}

func (w *Window) GetContainers() *fyne.Container {
	cont := container.NewVBox()

	for _, object := range w.Objects {
		cont.Add(object.Container)
	}

	return cont
}
