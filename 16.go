// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import (
	"fmt"
	"sort"
)
// В худшем случае (если выбран плохой опорный элемент): O(n^2)
// В лучшем случае (если опорный элемент всегда находится посередине массива): O(n * log(n))
//Рекурсивоное решение 
func quickSort(arr []int) []int { //
	if len(arr) < 2 {
		return arr
	}

	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	left := make([]int, 0, len(arr))
	right := make([]int, 0, len(arr))
	equal := make([]int, 0, len(arr))

	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v == pivot {
			equal = append(equal, v)
		} else {
			right = append(right, v)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	left = append(left, equal...)
	left = append(left, right...)

	return left
}



func main() {
	// Создаем срез с неотсортированными целыми числами
	nums := []int{4, 2, 7, 1, 9, 3, 5, 6, 8}

	sort.Ints(nums)
	fmt.Println(nums)

	
	nums = []int{4, 2, 7, 1, 9, 3, 5, 6, 8}

	nums = quickSort(nums)

	fmt.Println(nums)


}
