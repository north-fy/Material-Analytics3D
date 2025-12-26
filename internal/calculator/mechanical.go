package calculator

// Напряжения сжатия/растяжения q = FS, S перпендикулярна к силе F
// Относительная деформация e = |l-l2| / l, l2-l1 = delta l - изменение размера тела, l - начальная длина тела
// Модуль Юнга E = q / e
// и мб еще составление графика по Модулю Юнга, хз

type MechanicProperties struct {
	CompressiveStress   float64
	RelativeDeformation float64
	YoungsModulus       float64
}

type MechanicCalculator interface {
	Calculate() *MechanicProperties
	GetType() int
}

type Mech struct {
	CompressiveStress   CompressiveStress
	RelativeDeformation RelativeDeformation
	YoungsModulus       YoungsModulus
}

// CompressiveStress Напряжение сжатия/растяжения со значением силы F и площади попер. сеч. S, перпендикулярная к силе F
// Определенный тип - 0
type CompressiveStress struct {
	Force     float64
	CrossArea float64
}

func NewCompressiveStress(f, c float64) *CompressiveStress {
	return &CompressiveStress{Force: f, CrossArea: c}
}

func (c *CompressiveStress) Calculate() *MechanicProperties {
	compressiveStress := c.Force * c.CrossArea

	return &MechanicProperties{CompressiveStress: compressiveStress}
}

func (c *CompressiveStress) GetType() int {
	return CompressiveStressType
}

// RelativeDeformation Относительная деформация со значением длины L и длины растяжения/сжатия L2
// Определенный тип - 1
type RelativeDeformation struct {
	Lenght       float64
	ChangeLenght float64
}

func NewRelativeDeformation(l, l2 float64) *RelativeDeformation {
	return &RelativeDeformation{Lenght: l, ChangeLenght: l2}
}

func (r *RelativeDeformation) Calculate() *MechanicProperties {
	relativeDeformation := (r.ChangeLenght - r.Lenght) / r.Lenght

	return &MechanicProperties{RelativeDeformation: relativeDeformation}
}

func (r *RelativeDeformation) GetType() int {
	return RelativeDeformationType
}

// YoungsModulus Модуль Юнга со значением Напряжения сжатия/растяжения q и относительной деформации e
// Определенный тип - 2
type YoungsModulus struct {
	RelativeDeformationValue float64
	CompressiveStressValue   float64
}

func NewYoungsModulus(r, c float64) *YoungsModulus {
	return &YoungsModulus{RelativeDeformationValue: r, CompressiveStressValue: c}
}

func (y *YoungsModulus) Calculate() *MechanicProperties {
	youngsModulus := y.CompressiveStressValue / y.RelativeDeformationValue

	return &MechanicProperties{YoungsModulus: youngsModulus}
}

func (y *YoungsModulus) GetType() int {
	return YoungsModulusType
}

func CreateMechanicCalculator(mechanicType string, values map[string]float64) (MechanicCalculator, error) {
	switch mechanicType {
	case "CompressiveStress": // Напряжение растяжения/сжатия
		if len(values) == 2 {
			return NewCompressiveStress(values["Force"], values["CrossArea"]), nil
		}
		return nil, ErrEnoughArg

	case "RelativeDeformation": // Относительная деформация
		if len(values) == 2 {
			return NewRelativeDeformation(values["Lenght"], values["ChangeLenght"]), nil
		}
		return nil, ErrEnoughArg

	case "YoungsModulus": // Модуль Юнга
		if len(values) == 2 {
			return NewYoungsModulus(values["RelativeDeformationValue"], values["CompressiveStressValue"]), nil
		}
		return nil, ErrEnoughArg

	default: // Если тип неизвестен
		return nil, ErrUnknownType
	}
}
