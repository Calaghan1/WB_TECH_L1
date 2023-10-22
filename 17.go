// Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"fmt"
	"sort"
)

func main() {
	// Создаем срез с отсортированными целыми числами
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Выполняем бинарный поиск значения 5
	target := 5
	index := sort.SearchInts(nums, target)

	// Проверяем результат бинарного поиска
	if index < len(nums) && nums[index] == target {
		fmt.Printf("Значение %d найдено в позиции %d\n", target, index)
	} else {
		fmt.Printf("Значение %d не найдено\n", target)
	}
}