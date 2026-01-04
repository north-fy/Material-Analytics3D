package application

import (
	"fmt"
	"image/color"
	"strconv"
	"sync"
	"time"

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

	mu             *sync.Mutex
	visualSemaChan chan struct{}
}

func (r *Router) createBaseScreen() fyne.CanvasObject {
	// левые канвас объекты
	cs := CalcSettings{
		data:     make(map[string]float64),
		typeCalc: "",

		entryValues: make([]*widget.Entry, 4),
		content:     container.NewMax(),

		mu: &sync.Mutex{},
	}

	cs.labelDropDown1 = widget.NewLabel("Тип калькулятора:")
	cs.dropdown1 = widget.NewSelect(r.calcService.GetNames(), func(value string) {

		r.calcService.SetCurrentCalc(value)
		cs.dropdown2.Selected = ""

		cs.dropdown2.Options = r.calcService.GetNamesInterface()
		cs.dropdown2.Refresh()

		cs.visualSemaChan = make(chan struct{}, 1)
	})

	cs.dropdown1.PlaceHolder = "Выберите калькулятор"

	cs.separator1 = canvas.NewLine(color.Gray{0x80})
	cs.separator1.StrokeWidth = 2

	cs.labelDropDown2 = widget.NewLabel("Найти значение:")
	cs.dropdown2 = widget.NewSelect(r.calcService.GetNamesInterface(), func(value string) {
		for _, v := range cs.entryValues {
			v.Hide()
			v.Text = ""
		}

		values := r.calcService.GetNamesInterfaceValues(value)
		if values != nil {
			for i, v := range values {
				cs.entryValues[i].SetPlaceHolder(v)
				cs.entryValues[i].Show()

				cs.entryValues[i].Refresh()
			}
		}

		cs.typeCalc = value
	})
	cs.dropdown2.PlaceHolder = "Выберите значение"

	cs.separator2 = canvas.NewLine(color.Gray{0x80})
	cs.separator2.StrokeWidth = 2

	cs.labelAnswer = widget.NewLabel("Полученное значение:")

	for i := range 3 {
		cs.labelAnswerValue[i] = widget.NewLabel("")
	}

	leftPanel := container.NewVBox(
		cs.labelDropDown1,
		cs.dropdown1,
		cs.separator1,

		cs.labelDropDown2,
		cs.dropdown2,
		cs.separator2,

		cs.labelAnswer,
		cs.labelAnswerValue[0],
		cs.labelAnswerValue[1],
		cs.labelAnswerValue[2],
	)

	// верхние энтри поля
	cs.entryValues[0] = widget.NewEntry()
	cs.entryValues[0].Hide()

	cs.entryValues[1] = widget.NewEntry()
	cs.entryValues[1].Hide()

	cs.entryValues[2] = widget.NewEntry()
	cs.entryValues[2].Hide()

	cs.entryValues[3] = widget.NewEntry()
	cs.entryValues[3].Hide()

	topPanel := container.NewVBox(
		widget.NewLabel("Заполните информацию:"),
		container.NewGridWithColumns(2,
			cs.entryValues[0],
			cs.entryValues[1],
		),
		container.NewGridWithColumns(2,
			cs.entryValues[2],
			cs.entryValues[3],
		),
	)

	// нижние кнопки
	buttonAccept := widget.NewButton("Вычислить значение", func() {
		values := make(map[string]float64)

		for _, v := range cs.entryValues {
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
		answer, err := r.calcService.CalcCurrentCalc(cs.typeCalc, values)
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

			cs.labelAnswerValue[i].Text = text
			cs.labelAnswerValue[i].Refresh()
			i += 1
		}

		// history
		op := fmt.Sprintf("operation: %s | value: %v", cs.typeCalc, answer)
		addOperation(&cs.history, op)
	})
	buttonCancel := widget.NewButton("Сбросить", func() {
		for i := range 3 {
			cs.labelAnswerValue[i].Text = ""
			cs.labelAnswerValue[i].Refresh()
		}

		cs.dropdown1.Selected = ""
		cs.dropdown2.Selected = ""

		cs.dropdown1.Refresh()
		cs.dropdown2.Refresh()

		for _, v := range cs.entryValues {
			v.Hide()
		}

		if cs.visualSemaChan != nil {
			cs.visualSemaChan <- struct{}{}
		}
	})
	buttonOperations := widget.NewButton("Журнал операций", func() {
		var historyText string

		for _, v := range cs.history {
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
		if len(cs.content.Objects) > 0 {
			return
		}

		renderer := render.NewRenderer(750.0/1.5, 500.0/1.5)
		obj, err := render.GenerateObject(cs.typeCalc, render.GetColor(r.managerScreen.settings["color"]))
		if err != nil {
			return
		}
		visualObj := renderer.Render(obj)
		cs.content.Objects = []fyne.CanvasObject{visualObj}

		go func() {
			angle := 0.0
			for {
				select {
				case <-cs.visualSemaChan:
					cs.content.Objects = nil
					return
				default:
					angle += 0.01

					renderer.RotateY(obj, angle*0.2)

					fyne.Do(func() {
						if len(cs.content.Objects) > 0 {
							cs.content.Objects = []fyne.CanvasObject{renderer.Render(obj)}
							cs.content.Refresh()
						}
					})

					time.Sleep(100 * time.Millisecond)
				}
			}
		}()
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

	// правые кнопки
	logoRes1, _ := fyne.LoadResourceFromPath("./assets/whatsapp.png")
	icon1 := widget.NewIcon(logoRes1)

	logoRes2, _ := fyne.LoadResourceFromPath("./assets/Gopher.png")
	icon2 := widget.NewIcon(logoRes2)

	logoRes3, _ := fyne.LoadResourceFromPath("./assets/logo.png")
	icon3 := widget.NewIcon(logoRes3)

	buttonTop1 := widget.NewButton("set color", func() {
		r.managerScreen.setCurrentScreen("color")
	})
	buttonTop2 := widget.NewButton("version 1.0", func() {
		msg := fmt.Sprintf("version 1.0 \n Изменения: ...")
		dialog.ShowInformation("information", msg, r.managerScreen.window)
	})

	rightTopPanel := container.NewVBox(
		container.NewHBox(
			icon1,
			icon2,
			icon3,
		),
		buttonTop1,
		buttonTop2,
	)

	rightPanel := container.NewBorder(
		rightTopPanel,
		nil,
		nil,
		nil,
		nil,
	)

	// ЦЕНТР
	rect := canvas.NewRectangle(color.White)
	rect.StrokeColor = color.Gray{80}
	rect.StrokeWidth = 5.0

	rect.SetMinSize(fyne.NewSize(750.0/1.5, 500.0/1.5))

	rectContainer := container.NewStack(
		rect,
		cs.content,
	)

	rectContainer = container.NewCenter(rectContainer)

	// сборка
	mainContainer := container.NewBorder(
		topPanel,      // Верх
		bottomPanel,   // Низ
		leftPanel,     // Лево
		rightPanel,    // Право
		rectContainer, // Центр
	)

	return mainContainer
}

func addOperation(history *[10]string, newOp string) {
	for i := len(history) - 2; i >= 0; i-- {
		history[i+1] = history[i]
	}

	history[0] = newOp
}
