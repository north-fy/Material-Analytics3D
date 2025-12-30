package render

import "image/color"

// Цвета по умолчанию
var Colors = map[string]color.Color{
	"red":    color.RGBA{R: 255, G: 0, B: 0, A: 255},
	"green":  color.RGBA{R: 0, G: 255, B: 0, A: 255},
	"blue":   color.RGBA{R: 0, G: 0, B: 255, A: 255},
	"yellow": color.RGBA{R: 255, G: 255, B: 0, A: 255},
	"purple": color.RGBA{R: 128, G: 0, B: 128, A: 255},
	"cyan":   color.RGBA{R: 0, G: 255, B: 255, A: 255},
	"orange": color.RGBA{R: 255, G: 165, B: 0, A: 255},
	"pink":   color.RGBA{R: 255, G: 192, B: 203, A: 255},
	"gray":   color.RGBA{R: 128, G: 128, B: 128, A: 255},
}

// GetColor возвращает цвет по имени
func GetColor(name string) color.Color {
	if col, ok := Colors[name]; ok {
		return col
	}
	return Colors["gray"]
}
