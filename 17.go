// Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"fmt"
	"sort"
)


func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1 // Элемент не найден
}

func main() {
	// Создаем срез с отсортированными целыми числами
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Выполняем бинарный поиск значения 5
	target := 6
	index := sort.SearchInts(nums, target) // если не найдет возращает len(nums) + 1
	fmt.Println(index)
	// Проверяем результат бинарного поиска
	if index < len(nums) && nums[index] == target {
		fmt.Printf("Значение %d найдено в позиции %d\n", target, index)
	} else {
		fmt.Printf("Значение %d не найдено\n", target)
	}

	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index = binarySearch(nums, target)

	if index != -1 && nums[index] == target {
		fmt.Printf("Значение %d найдено в позиции %d\n", target, index)
	} else {
		fmt.Printf("Значение %d не найдено\n", target)
	}
}