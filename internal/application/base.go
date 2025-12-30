package application

import (
	"fmt"
	"image/color"
	"net/url"
	"strconv"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/north-fy/Material-Analytics3D/internal/render"
)

type CalcSettings struct {
	data     map[string]float64
	typeCalc string
	history  [10]string

	labelDropDown1 *widget.Label
	dropdown1      *widget.Select
	separator1     *canvas.Line

	labelDropDown2 *widget.Label
	dropdown2      *widget.Select
	separator2     *canvas.Line

	entryValues []*widget.Entry

	labelAnswer      *widget.Label
	labelAnswerValue [3]*widget.Label

	content *fyne.Container

	mu *sync.Mutex
}

func (r *Router) createBaseScreen() fyne.CanvasObject {
	// ================= ЛЕВАЯ ПАНЕЛЬ =================
	leftMenu := CalcSettings{
		data:        make(map[string]float64),
		entryValues: make([]*widget.Entry, 4),
		content:     container.NewMax(),

		mu: &sync.Mutex{},
	}

	leftMenu.labelDropDown1 = widget.NewLabel("Тип калькулятора:")
	leftMenu.dropdown1 = widget.NewSelect(r.calcService.GetNames(), func(value string) {

		r.calcService.SetCurrentCalc(value)
		leftMenu.dropdown2.Selected = ""

		leftMenu.dropdown2.Options = r.calcService.GetNamesInterface()
		leftMenu.dropdown2.Refresh()
	})

	leftMenu.dropdown1.PlaceHolder = "Выберите калькулятор"

	leftMenu.separator1 = canvas.NewLine(color.Gray{0x80})
	leftMenu.separator1.StrokeWidth = 2

	leftMenu.labelDropDown2 = widget.NewLabel("Найти значение:")
	leftMenu.dropdown2 = widget.NewSelect(r.calcService.GetNamesInterface(), func(value string) {
		for _, v := range leftMenu.entryValues {
			v.Hide()
			v.Text = ""
		}

		values := r.calcService.GetNamesInterfaceValues(value)
		if values != nil {
			for i, v := range values {
				leftMenu.entryValues[i].SetPlaceHolder(v)
				leftMenu.entryValues[i].Show()

				leftMenu.entryValues[i].Refresh()
			}
		}

		leftMenu.typeCalc = value
	})
	leftMenu.dropdown2.PlaceHolder = "Выберите значение"

	leftMenu.separator2 = canvas.NewLine(color.Gray{0x80})
	leftMenu.separator2.StrokeWidth = 2

	leftMenu.labelAnswer = widget.NewLabel("Полученное значение:")

	for i := range 3 {
		leftMenu.labelAnswerValue[i] = widget.NewLabel("")
	}

	leftPanel := container.NewVBox(
		leftMenu.labelDropDown1,
		leftMenu.dropdown1,
		leftMenu.separator1,

		leftMenu.labelDropDown2,
		leftMenu.dropdown2,
		leftMenu.separator2,

		leftMenu.labelAnswer,
		leftMenu.labelAnswerValue[0],
		leftMenu.labelAnswerValue[1],
		leftMenu.labelAnswerValue[2],
	)

	// ================= ВЕРХНЯЯ ПАНЕЛЬ =================
	leftMenu.entryValues[0] = widget.NewEntry()
	leftMenu.entryValues[0].Hide()

	leftMenu.entryValues[1] = widget.NewEntry()
	leftMenu.entryValues[1].Hide()

	leftMenu.entryValues[2] = widget.NewEntry()
	leftMenu.entryValues[2].Hide()

	leftMenu.entryValues[3] = widget.NewEntry()
	leftMenu.entryValues[3].Hide()

	// Контейнер для верхней панели
	topPanel := container.NewVBox(
		widget.NewLabel("Заполните информацию:"),
		container.NewGridWithColumns(2,
			leftMenu.entryValues[0],
			leftMenu.entryValues[1],
		),
		container.NewGridWithColumns(2,
			leftMenu.entryValues[2],
			leftMenu.entryValues[3],
		),
	)

	// ================= НИЖНЯЯ ПАНЕЛЬ =================

	buttonAccept := widget.NewButton("Вычислить значение", func() {
		values := make(map[string]float64)

		for _, v := range leftMenu.entryValues {
			if v.Text != "" {
				nameValue := v.PlaceHolder
				value, err := strconv.ParseFloat(v.Text, 64)
				if err != nil {
					dialog.ShowError(err, r.managerScreen.window)
				}

				values[nameValue] = value
			}
		}

		// мейби из-за канала вылетает после повторного вычисления
		answer, err := r.calcService.CalcCurrentCalc(leftMenu.typeCalc, values)
		if err != nil {
			dialog.ShowError(err, r.managerScreen.window)
		}

		var (
			i    int
			text string
		)

		for k, v := range answer {
			if v != 0 {
				text = fmt.Sprintf(" %s -> %.1f |", k, v)
			}

			leftMenu.labelAnswerValue[i].Text = text
			leftMenu.labelAnswerValue[i].Refresh()
			i += 1
		}

		// history
		op := fmt.Sprintf("operation: %s | value: %v", leftMenu.typeCalc, answer)
		addOperation(&leftMenu.history, op)
	})
	buttonCancel := widget.NewButton("Сбросить", func() {
		for i := range 3 {
			leftMenu.labelAnswerValue[i].Text = ""
			leftMenu.labelAnswerValue[i].Refresh()
		}

		leftMenu.dropdown1.Selected = ""
		leftMenu.dropdown2.Selected = ""

		leftMenu.dropdown1.Refresh()
		leftMenu.dropdown2.Refresh()

		for _, v := range leftMenu.entryValues {
			v.Hide()
		}
	})
	buttonOperations := widget.NewButton("Журнал операций", func() {
		var historyText string

		for _, v := range leftMenu.history {
			historyText += v + "\n"
		}

		dialog.ShowInformation("operations", historyText, r.managerScreen.window)
	})
	buttonNotation := widget.NewButton("Система счисления", func() {
		notion := fmt.Sprintf("Все значения записываются в исходных измерениях \n" +
			"Например: Масса - кг \n" +
			"Импульс - кг * м/с и так далее")
		dialog.ShowInformation("notion", notion, r.managerScreen.window)
	})
	buttonVisual := widget.NewButton("Визуализировать", func() {
		renderer := render.NewRenderer(750.0/1.5, 500.0/1.5)
		cube := render.CreateCube(2, render.GetColor("green"))
		cubeObj := renderer.Render(cube)

		leftMenu.content.Objects = []fyne.CanvasObject{cubeObj}

		//go func() {
		//	angle := 0.0
		//	for {
		//		angle += 0.02
		//
		//		renderer.Rotate(cube, angle*0.4, angle*0.6)
		//
		//		if len(leftMenu.content.Objects) > 0 {
		//			leftMenu.content.Objects = []fyne.CanvasObject{renderer.Render(cube)}
		//			leftMenu.content.Refresh()
		//		}
		//
		//		time.Sleep(16 * time.Millisecond)
		//	}
		//}()
	})

	bottomPanel := container.NewVBox(
		container.NewGridWithColumns(2,
			buttonAccept,
			buttonCancel,
		),
		container.NewGridWithColumns(3,
			buttonOperations,
			buttonNotation,
			buttonVisual,
		),
	)

	// ================= ПРАВАЯ ПАНЕЛЬ =================
	logoRes1, _ := fyne.LoadResourceFromPath("./assets/whatsapp.png")
	icon1 := widget.NewIcon(logoRes1)

	logoRes2, _ := fyne.LoadResourceFromPath("./assets/Gopher.png")
	icon2 := widget.NewIcon(logoRes2)

	logoRes3, _ := fyne.LoadResourceFromPath("./assets/logo.png")
	icon3 := widget.NewIcon(logoRes3)

	buttonTop1 := widget.NewHyperlink("github", parseURL("https://github.com/north-fy"))
	buttonTop2 := widget.NewHyperlink("telegram", parseURL("https://t.me/n0rth3am"))
	buttonTop3 := widget.NewButton("version 1.0", func() {
		msg := fmt.Sprintf("version 1.0 \n Изменения: ...")
		dialog.ShowInformation("information", msg, r.managerScreen.window)
	})

	// Контейнер для правой панели
	rightTopPanel := container.NewVBox(
		container.NewHBox(
			icon1,
			icon2,
			icon3,
		),
		buttonTop1,
		buttonTop2,
		buttonTop3,
	)

	rightPanel := container.NewBorder(
		rightTopPanel,
		nil,
		nil,
		nil,
		nil,
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
	rectContainer := container.NewStack(
		rect,
		leftMenu.content,
	)

	rectContainer = container.NewCenter(rectContainer)
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

func parseURL(urlStr string) *url.URL {
	link, _ := url.Parse(urlStr)
	return link
}

func addOperation(history *[10]string, newOp string) {
	for i := len(history) - 2; i >= 0; i-- {
		history[i+1] = history[i]
	}

	history[0] = newOp
}
