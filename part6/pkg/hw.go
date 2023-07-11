package hw

import (
	"log"
	"math"
)

// По условиям задачи, координаты не могут быть меньше 0.
type Geom struct {
	X1, Y1, X2, Y2 float64
}

// добавлен конструктор
func New(x1, y1, x2, y2 float64) *Geom {
	if x1 < 0 || x2 < 0 || y1 < 0 || y2 < 0 {
		log.Fatal("Координаты не могут быть меньше нуля")
	}
	geom := Geom{X1: x1, Y1: y1, X2: x2, Y2: y2}
	return &geom
}

// добавлена функция типа Geom
func (g *Geom) CalcDist() (distance float64) {
	return math.Sqrt(math.Pow(g.X2-g.X1, 2) + math.Pow(g.Y2-g.Y1, 2))
}
