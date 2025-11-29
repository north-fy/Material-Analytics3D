package calculator

// Количество теплоты Q = cm delta T (t2-t1)
// Тепловое расширение delta L = a * L * delta T (t2 - t1), a - коэффициент линейного теплового рассширения

type ThermicProperties struct {
	AmountHeat       float64
	ThermalExpansion float64
}

type ThermicCalculator interface {
	Calculate() *ThermicProperties
	GetType() int
}

type Therm struct {
	AmountHeat
	ThermalExpansion
}

// AmountHeat Количество теплоты со значениями удельной теплоемкости c, массы m и изминения температуры delta T
// Определенный тип - 0
type AmountHeat struct {
	SpecificHeat float64
	Mass         float64
	Temperature1 float64
	Temperature2 float64
}

func NewAmountHeat(s, m, t1, t2 float64) *AmountHeat {
	return &AmountHeat{SpecificHeat: s, Mass: m, Temperature1: t1, Temperature2: t2}
}

func (a *AmountHeat) Calculate() *ThermicProperties {
	amountHeat := a.SpecificHeat * a.Mass * (a.Temperature2 - a.Temperature1)

	return &ThermicProperties{AmountHeat: amountHeat}
}

func (a *AmountHeat) GetType() int {
	return AmountHeatType
}

// ThermalExpansion Тепловое расширение со значениями коэффициента тепл. расш a, длины объекта L,
// И изменение температуры delta T
// Определенный тип - 1
type ThermalExpansion struct {
	Coefficient  float64
	Lenght       float64
	Temperature1 float64
	Temperature2 float64
}

func NewThermalExpansion(c, l, t1, t2 float64) *ThermalExpansion {
	return &ThermalExpansion{Coefficient: c, Lenght: l, Temperature1: t1, Temperature2: t2}
}

func (t *ThermalExpansion) Calculate() *ThermicProperties {
	thermalExpansion := t.Coefficient * t.Lenght * (t.Temperature2 - t.Temperature1)

	return &ThermicProperties{ThermalExpansion: thermalExpansion}
}

func (t *ThermalExpansion) GetType() int {
	return ThermalExpansionType
}

func CreateThermicCalculator(thermicType int, values ...float64) (ThermicCalculator, error) {
	switch thermicType {
	case 0: // Количество теплоты
		if len(values) == 4 {
			return NewAmountHeat(values[0], values[1], values[2], values[3]), nil
		}
		return nil, ErrEnoughArg

	case 1: // Тепловое расширение
		if len(values) == 4 {
			return NewThermalExpansion(values[0], values[1], values[2], values[3]), nil
		}
		return nil, ErrEnoughArg

	default:
		return nil, ErrUnknownType
	}
}
