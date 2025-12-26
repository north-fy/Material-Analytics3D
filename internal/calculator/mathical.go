package calculator

import (
	"math"
)

type MathProperties struct {
	TotalArea float64
	Volume    float64
}

type MathCalculator interface {
	Calculate() *MathProperties
	GetType() int
}

type Math struct {
	Cube           Cube
	Parallelepiped Parallelepiped
	Pyramid        Pyramid
}

// Cube объект со значением длины одной стороны A
// Определенный тип - 0
type Cube struct {
	Side float64
}

func NewCube(side float64) *Cube {
	return &Cube{Side: side}
}

func (c *Cube) Calculate() *MathProperties {
	totalArea := 6 * math.Pow(c.Side, 2)
	volume := math.Pow(c.Side, 3)

	return &MathProperties{TotalArea: totalArea, Volume: volume}
}

func (c *Cube) GetType() int {
	return CubeType
}

// Parallelepiped объект со значением длины A, ширины B и высоты C
// Определенный тип - 1
type Parallelepiped struct {
	Lenght float64
	Width  float64
	Height float64
}

func NewParallelepiped(a, b, c float64) *Parallelepiped {
	return &Parallelepiped{Lenght: a, Width: b, Height: c}
}

func (p *Parallelepiped) Calculate() *MathProperties {
	totalArea := (p.Lenght + p.Width) * 2 * p.Height
	volume := p.Lenght * p.Width * p.Height

	return &MathProperties{TotalArea: totalArea, Volume: volume}
}

func (p *Parallelepiped) GetType() int {
	return ParallelepipedType
}

// Pyramid объект со значением длины квадрата основания пирамиды A, высоты от центра основания H
// Определенный тип - 2
type Pyramid struct {
	BaseSide float64
	Height   float64
}

func NewPyramid(a, h float64) *Pyramid {
	return &Pyramid{BaseSide: a, Height: h}
}

func (p *Pyramid) Calculate() *MathProperties {
	totalArea := (2 * p.BaseSide * p.Height) + math.Pow(p.BaseSide, 2)
	volume := (1 / 3) * math.Pow(p.BaseSide, 2) * p.Height

	return &MathProperties{TotalArea: totalArea, Volume: volume}
}

func (p *Pyramid) GetType() int {
	return PyramidType
}

// CreateMathCalculator создает фигуру по типу фигуры и значениям
// Выдает значения в метрах, стоит учитывать
func CreateMathCalculator(ShapeType string, values map[string]float64) (MathCalculator, error) {
	switch ShapeType {
	case "Cube": // Cube
		if len(values) == 1 {
			return NewCube(values["Side"]), nil
		}
		return nil, ErrEnoughArg

	case "Parallelepiped": // Parallelepiped
		if len(values) == 3 {
			return NewParallelepiped(values["Lenght"], values["Width"], values["Height"]), nil
		}
		return nil, ErrEnoughArg

	case "Pyramid": // Pyramid
		if len(values) == 2 {
			return NewPyramid(values["BaseSide"], values["Height"]), nil
		}
		return nil, ErrEnoughArg

	default: // Если тип неизвестен
		return nil, ErrUnknownType
	}
}
