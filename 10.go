// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
package main

import (
	"fmt"
	"sort"
)

func Exercise_10() {
	groups := make(map[int][]float64)
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	for _, i := range temperatures {
		groups[int(i/10) * 10] = append(groups[int(i/10)*10], i)
	}
	for group, tmp := range groups {
		sort.Float64s(tmp)
		fmt.Printf("%d градусов группа: %v\n", group, tmp)
	}
}