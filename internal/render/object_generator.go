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
	Indices  []int // Индексы вершин в общем массиве
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

// CreatePyramid создает пирамиду с правильными индексами
func CreatePyramid(baseSize, height float64, col color.Color) *Mesh {
	half := baseSize / 2

	vertices := []Vector3{
		// Основание
		{half, -half, -half},  // 0
		{-half, -half, -half}, // 1
		{-half, -half, half},  // 2
		{half, -half, half},   // 3
		// Вершина
		{0, -height, 0}, // 4 (смещена вниз)
	}

	faces := []Face{
		// Основание (квадрат)
		{
			Vertices: []Vector3{vertices[0], vertices[1], vertices[2], vertices[3]},
			Indices:  []int{0, 1, 2, 3},
			Color:    adjustColor(col, -20),
		},
		// Боковые грани (треугольники)
		{
			Vertices: []Vector3{vertices[0], vertices[1], vertices[4]},
			Indices:  []int{0, 1, 4},
			Color:    adjustColor(col, -30),
		},
		{
			Vertices: []Vector3{vertices[1], vertices[2], vertices[4]},
			Indices:  []int{1, 2, 4},
			Color:    adjustColor(col, -20),
		},
		{
			Vertices: []Vector3{vertices[2], vertices[3], vertices[4]},
			Indices:  []int{2, 3, 4},
			Color:    adjustColor(col, -10),
		},
		{
			Vertices: []Vector3{vertices[3], vertices[0], vertices[4]},
			Indices:  []int{3, 0, 4},
			Color:    adjustColor(col, -40),
		},
	}

	return &Mesh{Vertices: vertices, Faces: faces}
}

func CreateParallelepiped(width, height, depth float64, col color.Color) *Mesh {
	w, h, d := width, height, depth

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
		// Передняя
		{
			Vertices: []Vector3{vertices[0], vertices[1], vertices[2], vertices[3]},
			Indices:  []int{0, 1, 2, 3},
			Color:    col,
		},
		// Задняя
		{
			Vertices: []Vector3{vertices[4], vertices[5], vertices[6], vertices[7]},
			Indices:  []int{4, 5, 6, 7},
			Color:    adjustColor(col, -40),
		},
		// Верхняя
		{
			Vertices: []Vector3{vertices[3], vertices[2], vertices[6], vertices[7]},
			Indices:  []int{3, 2, 6, 7},
			Color:    adjustColor(col, 20),
		},
		// Нижняя
		{
			Vertices: []Vector3{vertices[0], vertices[1], vertices[5], vertices[4]},
			Indices:  []int{0, 1, 5, 4},
			Color:    adjustColor(col, -20),
		},
		// Левая
		{
			Vertices: []Vector3{vertices[4], vertices[0], vertices[3], vertices[7]},
			Indices:  []int{4, 0, 3, 7},
			Color:    adjustColor(col, -10),
		},
		// Правая
		{
			Vertices: []Vector3{vertices[1], vertices[5], vertices[6], vertices[2]},
			Indices:  []int{1, 5, 6, 2},
			Color:    adjustColor(col, 10),
		},
	}

	return &Mesh{Vertices: vertices, Faces: faces}
}

//// CreateCube создает куб с правильными индексами
//func CreateCube(size float64, col color.Color) *Mesh {
//	half := size / 2
//
//	vertices := []Vector3{
//		// Передняя грань
//		{-half, -half, half}, // 0
//		{half, -half, half},  // 1
//		{half, half, half},   // 2
//		{-half, half, half},  // 3
//		// Задняя грань
//		{-half, -half, -half}, // 4
//		{half, -half, -half},  // 5
//		{half, half, -half},   // 6
//		{-half, half, -half},  // 7
//	}
//
//	faces := []Face{
//		// Передняя
//		{
//			Vertices: []Vector3{vertices[0], vertices[1], vertices[2], vertices[3]},
//			Indices:  []int{0, 1, 2, 3},
//			Color:    col,
//		},
//		// Задняя
//		{
//			Vertices: []Vector3{vertices[4], vertices[5], vertices[6], vertices[7]},
//			Indices:  []int{4, 5, 6, 7},
//			Color:    adjustColor(col, -40),
//		},
//		// Верхняя
//		{
//			Vertices: []Vector3{vertices[3], vertices[2], vertices[6], vertices[7]},
//			Indices:  []int{3, 2, 6, 7},
//			Color:    adjustColor(col, 20),
//		},
//		// Нижняя
//		{
//			Vertices: []Vector3{vertices[0], vertices[1], vertices[5], vertices[4]},
//			Indices:  []int{0, 1, 5, 4},
//			Color:    adjustColor(col, -20),
//		},
//		// Левая
//		{
//			Vertices: []Vector3{vertices[4], vertices[0], vertices[3], vertices[7]},
//			Indices:  []int{4, 0, 3, 7},
//			Color:    adjustColor(col, -10),
//		},
//		// Правая
//		{
//			Vertices: []Vector3{vertices[1], vertices[5], vertices[6], vertices[2]},
//			Indices:  []int{1, 5, 6, 2},
//			Color:    adjustColor(col, 10),
//		},
//	}
//
//	return &Mesh{Vertices: vertices, Faces: faces}
//}

// Render рисует объект
func (r *Renderer) Render(mesh *Mesh) fyne.CanvasObject {
	content := container.NewWithoutLayout()

	// Сначала рисуем заполненные грани
	for _, face := range mesh.Faces {
		if len(face.Vertices) < 3 {
			continue
		}
		r.drawFilledFace(content, face)
	}

	// Затем рисуем контуры поверх
	for _, face := range mesh.Faces {
		if len(face.Vertices) < 2 {
			continue
		}
		r.drawFaceOutline(content, face)
	}

	return content
}

// drawFaceOutline рисует контур грани
func (r *Renderer) drawFaceOutline(container *fyne.Container, face Face) {
	for i := 0; i < len(face.Vertices); i++ {
		next := (i + 1) % len(face.Vertices)

		x1, y1 := r.project(face.Vertices[i])
		x2, y2 := r.project(face.Vertices[next])

		line := canvas.NewLine(face.Color)
		line.Position1 = fyne.NewPos(float32(x1), float32(y1))
		line.Position2 = fyne.NewPos(float32(x2), float32(y2))
		line.StrokeWidth = 2

		container.Add(line)
	}
}

// drawFilledFace рисует заполненную грань
func (r *Renderer) drawFilledFace(container *fyne.Container, face Face) {
	// Создаем полигон для заполнения
	points := make([]fyne.Position, len(face.Vertices))
	for i, v := range face.Vertices {
		x, y := r.project(v)
		points[i] = fyne.NewPos(float32(x), float32(y))
	}

	// Создаем и добавляем полигон
	//polygon := canvas.NewPolygon(, face.Color)
	//container.Add(polygon)
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

// RotateY вращает объект вокруг Y
func (r *Renderer) RotateY(mesh *Mesh, angle float64) {
	cosA := math.Cos(angle)
	sinA := math.Sin(angle)

	// Вращаем все вершины
	for i := range mesh.Vertices {
		v := &mesh.Vertices[i]

		// Вращение вокруг оси Y
		x := v.X*cosA + v.Z*sinA
		z := -v.X*sinA + v.Z*cosA

		v.X, v.Z = x, z
	}

	// ОБНОВЛЯЕМ ВЕРШИНЫ В ГРАНЯХ, используя индексы
	for faceIdx := range mesh.Faces {
		for vertIdx, vertexIdx := range mesh.Faces[faceIdx].Indices {
			if vertexIdx < len(mesh.Vertices) {
				mesh.Faces[faceIdx].Vertices[vertIdx] = mesh.Vertices[vertexIdx]
			}
		}
	}
}

// RotateX вращает объект вокруг X
func (r *Renderer) RotateX(mesh *Mesh, angle float64) {
	cosA := math.Cos(angle)
	sinA := math.Sin(angle)

	// Вращаем все вершины
	for i := range mesh.Vertices {
		v := &mesh.Vertices[i]

		// Вращение вокруг оси X
		y := v.Y*cosA - v.Z*sinA
		z := v.Y*sinA + v.Z*cosA

		v.Y, v.Z = y, z
	}

	// Обновляем вершины в гранях
	for faceIdx := range mesh.Faces {
		for vertIdx, vertexIdx := range mesh.Faces[faceIdx].Indices {
			if vertexIdx < len(mesh.Vertices) {
				mesh.Faces[faceIdx].Vertices[vertIdx] = mesh.Vertices[vertexIdx]
			}
		}
	}
}

// RotateZ вращает объект вокруг Z
func (r *Renderer) RotateZ(mesh *Mesh, angle float64) {
	cosA := math.Cos(angle)
	sinA := math.Sin(angle)

	// Вращаем все вершины
	for i := range mesh.Vertices {
		v := &mesh.Vertices[i]

		// Вращение вокруг оси Z
		x := v.X*cosA - v.Y*sinA
		y := v.X*sinA + v.Y*cosA

		v.X, v.Y = x, y
	}

	// Обновляем вершины в гранях
	for faceIdx := range mesh.Faces {
		for vertIdx, vertexIdx := range mesh.Faces[faceIdx].Indices {
			if vertexIdx < len(mesh.Vertices) {
				mesh.Faces[faceIdx].Vertices[vertIdx] = mesh.Vertices[vertexIdx]
			}
		}
	}
}

// Обновленная функция для использования углов в градусах
func (r *Renderer) RotateDegrees(mesh *Mesh, angleX, angleY, angleZ float64) {
	// Конвертируем градусы в радианы
	radX := angleX * math.Pi / 180
	radY := angleY * math.Pi / 180
	radZ := angleZ * math.Pi / 180

	// Вращаем по всем осям
	if radX != 0 {
		r.RotateX(mesh, radX)
	}
	if radY != 0 {
		r.RotateY(mesh, radY)
	}
	if radZ != 0 {
		r.RotateZ(mesh, radZ)
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
