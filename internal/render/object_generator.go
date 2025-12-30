package render

import (
	"fmt"
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// Vector3 - 3D точка
type Vector3 struct {
	X, Y, Z float64
}

// Face - грань объекта
type Face struct {
	Vertices []Vector3
	Color    color.Color
}

// Mesh - 3D объект
type Mesh struct {
	Vertices []Vector3
	Faces    []Face
}

// Renderer - рендерер 3D объектов
type Renderer struct {
	Scale            float64
	CenterX, CenterY float64
	CameraZ          float64
}

// NewRenderer создает новый рендерер
func NewRenderer(width, height float64) *Renderer {
	return &Renderer{
		Scale:   80,
		CenterX: width / 2,
		CenterY: height / 2,
		CameraZ: 5,
	}
}

// CreatePyramid создает пирамиду
func CreatePyramid(baseSize, height float64, col color.Color) *Mesh {
	half := baseSize / 2

	vertices := []Vector3{
		// Основание
		{-half, -half, -half}, // 0
		{half, -half, -half},  // 1
		{half, -half, half},   // 2
		{-half, -half, half},  // 3
		// Вершина
		{0, height, 0}, // 4
	}

	faces := []Face{
		// Основание (4 стороны)
		{
			Vertices: []Vector3{vertices[0], vertices[1], vertices[2], vertices[3]},
			Color:    col,
		},
		// Боковые грани (треугольники - 3 стороны)
		{
			Vertices: []Vector3{vertices[0], vertices[1], vertices[4]},
			Color:    adjustColor(col, -30),
		},
		{
			Vertices: []Vector3{vertices[1], vertices[2], vertices[4]},
			Color:    adjustColor(col, -20),
		},
		{
			Vertices: []Vector3{vertices[2], vertices[3], vertices[4]},
			Color:    adjustColor(col, -10),
		},
		{
			Vertices: []Vector3{vertices[3], vertices[0], vertices[4]},
			Color:    adjustColor(col, -40),
		},
	}

	return &Mesh{Vertices: vertices, Faces: faces}
}

// CreateCube создает куб
func CreateCube(size float64, col color.Color) *Mesh {
	half := size / 2

	vertices := []Vector3{
		// Передняя грань
		{-half, -half, half}, // 0
		{half, -half, half},  // 1
		{half, half, half},   // 2
		{-half, half, half},  // 3
		// Задняя грань
		{-half, -half, -half}, // 4
		{half, -half, -half},  // 5
		{half, half, -half},   // 6
		{-half, half, -half},  // 7
	}

	faces := []Face{
		// Передняя (4 стороны)
		{
			Vertices: []Vector3{vertices[0], vertices[1], vertices[2], vertices[3]},
			Color:    col,
		},
		// Задняя (4 стороны)
		{
			Vertices: []Vector3{vertices[4], vertices[5], vertices[6], vertices[7]},
			Color:    adjustColor(col, -40),
		},
		// Верхняя (4 стороны)
		{
			Vertices: []Vector3{vertices[3], vertices[2], vertices[6], vertices[7]},
			Color:    adjustColor(col, 20),
		},
		// Нижняя (4 стороны)
		{
			Vertices: []Vector3{vertices[0], vertices[1], vertices[5], vertices[4]},
			Color:    adjustColor(col, -20),
		},
		// Левая (4 стороны)
		{
			Vertices: []Vector3{vertices[4], vertices[0], vertices[3], vertices[7]},
			Color:    adjustColor(col, -10),
		},
		// Правая (4 стороны)
		{
			Vertices: []Vector3{vertices[1], vertices[5], vertices[6], vertices[2]},
			Color:    adjustColor(col, 10),
		},
	}

	return &Mesh{Vertices: vertices, Faces: faces}
}

// CreateParallelepiped создает параллелепипед
func CreateParallelepiped(width, height, depth float64, col color.Color, stretch string) *Mesh {
	w, h, d := width, height, depth

	switch stretch {
	case "width":
		w *= 1.5
	case "height":
		h *= 1.5
	case "depth":
		d *= 1.5
	}

	halfW := w / 2
	halfH := h / 2
	halfD := d / 2

	vertices := []Vector3{
		// Передняя грань
		{-halfW, -halfH, halfD}, // 0
		{halfW, -halfH, halfD},  // 1
		{halfW, halfH, halfD},   // 2
		{-halfW, halfH, halfD},  // 3
		// Задняя грань
		{-halfW, -halfH, -halfD}, // 4
		{halfW, -halfH, -halfD},  // 5
		{halfW, halfH, -halfD},   // 6
		{-halfW, halfH, -halfD},  // 7
	}

	faces := []Face{
		// Все грани по 4 стороны
		{
			Vertices: []Vector3{vertices[0], vertices[1], vertices[2], vertices[3]},
			Color:    col,
		},
		{
			Vertices: []Vector3{vertices[4], vertices[5], vertices[6], vertices[7]},
			Color:    adjustColor(col, -40),
		},
		{
			Vertices: []Vector3{vertices[3], vertices[2], vertices[6], vertices[7]},
			Color:    adjustColor(col, 20),
		},
		{
			Vertices: []Vector3{vertices[0], vertices[1], vertices[5], vertices[4]},
			Color:    adjustColor(col, -20),
		},
		{
			Vertices: []Vector3{vertices[4], vertices[0], vertices[3], vertices[7]},
			Color:    adjustColor(col, -10),
		},
		{
			Vertices: []Vector3{vertices[1], vertices[5], vertices[6], vertices[2]},
			Color:    adjustColor(col, 10),
		},
	}

	return &Mesh{Vertices: vertices, Faces: faces}
}

// Render рисует объект
func (r *Renderer) Render(mesh *Mesh) fyne.CanvasObject {
	content := container.NewWithoutLayout()

	for _, face := range mesh.Faces {
		if len(face.Vertices) < 3 {
			continue
		}

		// Создаем линии для отрисовки граней
		r.drawFace(content, face)
	}

	return content
}

// drawFace рисует одну грань
func (r *Renderer) drawFace(container *fyne.Container, face Face) {
	// Рисуем линии между всеми вершинами грани
	for i := 0; i < len(face.Vertices); i++ {
		next := (i + 1) % len(face.Vertices)

		x1, y1 := r.project(face.Vertices[i])
		x2, y2 := r.project(face.Vertices[next])

		line := canvas.NewLine(face.Color)
		line.Position1 = fyne.NewPos(float32(x1), float32(y1))
		line.Position2 = fyne.NewPos(float32(x2), float32(y2))
		line.StrokeWidth = 3

		container.Add(line)
	}

	// Дополнительно: рисуем заполнение для грани
	if len(face.Vertices) >= 3 {
		r.drawFilledFace(container, face)
	}
}

// drawFilledFace рисует заполненную грань
func (r *Renderer) drawFilledFace(container *fyne.Container, face Face) {
	// Определяем количество сторон для полигона
	numSides := len(face.Vertices)
	if numSides < 3 {
		numSides = 3
	}

	// Создаем полигон с правильным количеством сторон
	polygon := canvas.NewPolygon(uint(numSides), face.Color)

	// Находим границы грани
	minX, minY := math.MaxFloat64, math.MaxFloat64
	maxX, maxY := -math.MaxFloat64, -math.MaxFloat64

	for _, v := range face.Vertices {
		x, y := r.project(v)
		if x < minX {
			minX = x
		}
		if y < minY {
			minY = y
		}
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	// Устанавливаем позицию и размер полигона
	pos := fyne.NewPos(float32(minX), float32(minY))
	size := fyne.NewSize(float32(maxX-minX), float32(maxY-minY))

	// Используем Move и Resize
	polygon.Move(pos)
	polygon.Resize(size)

	// Делаем полигон полупрозрачным для видимости линий
	if rgba, ok := face.Color.(color.RGBA); ok {
		rgba.A = 100 // Полупрозрачный
		polygon.FillColor = rgba
	}

	container.Add(polygon)
}

// project преобразует 3D в 2D
func (r *Renderer) project(v Vector3) (float64, float64) {
	// Простая перспективная проекция
	factor := r.CameraZ / (r.CameraZ + v.Z)

	x := v.X * factor * r.Scale
	y := v.Y * factor * r.Scale

	// Центрирование
	x += r.CenterX
	y += r.CenterY

	return x, y
}

// Rotate вращает объект
func (r *Renderer) Rotate(mesh *Mesh, angleX, angleY float64) {
	cosX, sinX := math.Cos(angleX), math.Sin(angleX)
	cosY, sinY := math.Cos(angleY), math.Sin(angleY)

	// Вращаем все вершины
	for i := range mesh.Vertices {
		v := &mesh.Vertices[i]

		// Вращение вокруг Y
		x := v.X*cosY - v.Z*sinY
		z := v.X*sinY + v.Z*cosY
		v.X, v.Z = x, z

		// Вращение вокруг X
		y := v.Y*cosX - v.Z*sinX
		z = v.Y*sinX + v.Z*cosX
		v.Y, v.Z = y, z
	}

	// Обновляем грани
	for i := range mesh.Faces {
		for j := range mesh.Faces[i].Vertices {
			// Используем соответствующие вершины
			mesh.Faces[i].Vertices[j] = mesh.Vertices[j]
		}
	}
}

// adjustColor регулирует яркость
func adjustColor(c color.Color, delta int) color.Color {
	r, g, b, a := c.RGBA()

	clamp := func(x int) uint8 {
		if x < 0 {
			return 0
		}
		if x > 255 {
			return 255
		}
		return uint8(x)
	}

	r8 := clamp(int(r>>8) + delta)
	g8 := clamp(int(g>>8) + delta)
	b8 := clamp(int(b>>8) + delta)

	return color.RGBA{R: r8, G: g8, B: b8, A: uint8(a >> 8)}
}

// String для отладки
func (v Vector3) String() string {
	return fmt.Sprintf("(%.2f, %.2f, %.2f)", v.X, v.Y, v.Z)
}
