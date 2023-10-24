// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

// Структура Point представляет точку в двумерном пространстве.
type Point struct {
	x, y float64
}

// Конструктор для создания новой точки.
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Метод для вычисления расстояния между текущей точкой и другой точкой p.
func (p Point) DistanceTo(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)// Находим гипотнезу 
}

func main() {
	// Создание двух точек
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	// Вычисление расстояния между точками
	distance := point1.DistanceTo(point2)

	// Вывод результата
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}