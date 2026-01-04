package render

import (
	"errors"
	"image/color"
)

var (
	size = 2.0
)

func GenerateObject(name string, col color.Color) (*Mesh, error) {
	switch name {
	case "Cube":
		return CreateParallelepiped(size, size, size, col), nil
	case "Parallelepiped":
		return CreateParallelepiped(4.0, size, size, col), nil
	case "Pyramid":
		return CreatePyramid(-1, 1, col), nil
	}

	return nil, errors.New("this object not correct")
}
