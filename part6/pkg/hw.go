package hw

import (
	"errors"
	"math"
)

// Убрана структура Geom
// Функция реализована без структуры
// Поправлен тест без структуры

func CalculateDistance(X1, Y1, X2, Y2 float64) (float64, error) {
	if X1 < 0 || X2 < 0 || Y1 < 0 || Y2 < 0 {
		return -1, errors.New("Координаты не могут быть меньше нуля")
	}

	return math.Sqrt(math.Pow(X2-X1, 2) + math.Pow(Y2-Y1, 2)), nil
}
