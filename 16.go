// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import (
	"fmt"
	"sort"
)

func main() {
	// Создаем срез с неотсортированными целыми числами
	nums := []int{4, 2, 7, 1, 9, 3, 5, 6, 8}

	// Используем функцию sort.Ints для сортировки среза
	sort.Ints(nums)
	// func Sort(data Interface) {
	// 	n := data.Len()
	// 	if n <= 1 {
	// 		return
	// 	}
	// 	limit := bits.Len(uint(n))
	// 	pdqsort(data, 0, n, limit) quicksort в го
	// }  
	// Выводим отсортированный срез
	fmt.Println(nums)
}