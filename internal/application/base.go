package application

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (r *Router) createBaseScreen() fyne.CanvasObject {
	// ================= ЛЕВАЯ ПАНЕЛЬ =================
	// Первое выпадающее меню
	labelDropDown1 := widget.NewLabel("Меню 1:")
	dropdown1 := widget.NewSelect(r.calcService.GetNames(), func(value string) {
		r.calcService.CurrentCalc = r.calcService.SpecificCalc[value]
	})
	dropdown1.PlaceHolder = "Выберите калькулятор"
	// Разделительная линия
	separator1 := canvas.NewLine(color.Gray{0x80})
	separator1.StrokeWidth = 1

	// Второе выпадающее меню
	labelDropDown2 := widget.NewLabel("Меню 2:")
	dropdown2 := widget.NewSelect(r.calcService.GetNamesInterface(), func(value string) {
		return
	})
	dropdown2.PlaceHolder = "Выберите значение"

	// Разделительная линия
	separator2 := canvas.NewLine(color.Gray{0x80})
	separator2.StrokeWidth = 1

	// Контейнер для левой панели
	leftPanel := container.NewVBox(
		labelDropDown1,
		dropdown1,
		separator1,
		labelDropDown2,
		dropdown2,
		separator2,
	)

	// ================= ВЕРХНЯЯ ПАНЕЛЬ =================
	// Поля ввода
	entry1 := widget.NewEntry()
	entry1.SetPlaceHolder("Поле ввода 1")

	entry2 := widget.NewEntry()
	entry2.SetPlaceHolder("Поле ввода 2")

	entry3 := widget.NewEntry()
	entry3.SetPlaceHolder("Поле ввода 3")

	entry4 := widget.NewEntry()
	entry4.SetPlaceHolder("Поле ввода 4")

	entry5 := widget.NewEntry()
	entry5.SetPlaceHolder("Поле ввода 5")

	// Контейнер для верхней панели
	topPanel := container.NewVBox(
		widget.NewLabel("Заполните информацию:"),
		container.NewGridWithColumns(3,
			entry1,
			entry2,
			entry3,
		),
		container.NewGridWithColumns(2,
			entry4,
			entry5,
		),
	)

	// ================= НИЖНЯЯ ПАНЕЛЬ =================
	// Чекбоксы
	check1 := widget.NewCheck("Опция A", func(checked bool) {})
	check2 := widget.NewCheck("Опция B", func(checked bool) {})
	check3 := widget.NewCheck("Опция C", func(checked bool) {})
	check4 := widget.NewCheck("Опция D", func(checked bool) {})
	check5 := widget.NewCheck("Опция E", func(checked bool) {})

	// Контейнер для нижней панели
	bottomPanel := container.NewVBox(
		widget.NewLabel("Выберите опции:"),
		container.NewGridWithColumns(3,
			check1,
			check2,
			check3,
		),
		container.NewGridWithColumns(2,
			check4,
			check5,
		),
	)

	// ================= ПРАВАЯ ПАНЕЛЬ =================
	// Верхняя часть правой панели
	// Круглые иконки
	icon1 := widget.NewIcon(theme.InfoIcon())
	icon2 := widget.NewIcon(theme.WarningIcon())

	// Кнопки для верхней части
	buttonTop1 := widget.NewButton("Кнопка 1", func() {})
	buttonTop2 := widget.NewButton("Кнопка 2", func() {})
	buttonTop3 := widget.NewButton("Кнопка 3", func() {})

	// Центральная и нижняя части правой панели
	buttonCenter1 := widget.NewButton("Центр 1", func() {})
	buttonCenter2 := widget.NewButton("Центр 2", func() {})
	buttonBottom1 := widget.NewButton("Низ 1", func() {})
	buttonBottom2 := widget.NewButton("Низ 2", func() {})
	buttonBottom3 := widget.NewButton("Низ 3", func() {})

	// Контейнер для правой панели
	rightTopPanel := container.NewVBox(
		container.NewHBox(
			icon1,
			icon2,
		),
		buttonTop1,
		buttonTop2,
		buttonTop3,
	)

	rightCenterPanel := container.NewVBox(
		buttonCenter1,
		buttonCenter2,
	)

	rightBottomPanel := container.NewVBox(
		buttonBottom1,
		buttonBottom2,
		buttonBottom3,
	)

	rightPanel := container.NewBorder(
		rightTopPanel,
		rightBottomPanel,
		nil,
		nil,
		rightCenterPanel,
	)

	// ================= ЦЕНТРАЛЬНЫЙ ПРЯМОУГОЛЬНИК =================
	// Создаем белый прямоугольник с серой обводкой
	rect := canvas.NewRectangle(color.White)
	rect.StrokeColor = color.Gray{80}
	rect.StrokeWidth = 5.0

	// Размеры прямоугольника вычисляются относительно размеров окна
	// Ширина и высота прямоугольника будут составлять 1/3 от размеров окна
	// с отступами по 1/3 с каждой стороны
	rect.SetMinSize(fyne.NewSize(750.0/1.5, 500.0/1.5))

	// Контейнер для прямоугольника (для центрирования)
	rectContainer := container.NewCenter(rect)

	// ================= ОСНОВНОЙ МАКЕТ =================
	// Создаем основной макет с помощью Border
	// Сверху - верхняя панель
	// Снизу - нижняя панель
	// Слева - левая панель
	// Справа - правая панель
	// В центре - прямоугольник
	mainContainer := container.NewBorder(
		topPanel,      // Верх
		bottomPanel,   // Низ
		leftPanel,     // Лево
		rightPanel,    // Право
		rectContainer, // Центр
	)

	return mainContainer
}
