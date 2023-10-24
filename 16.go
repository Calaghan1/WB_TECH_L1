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



// func partition(arr []int, low, high int) int {
// 	pivot := arr[high]
// 	i := low - 1

// 	for j := low; j <= high-1; j++ {
// 		if arr[j] < pivot {
// 			i++
// 			arr[i], arr[j] = arr[j], arr[i]
// 		}
// 	}

// 	arr[i+1], arr[high] = arr[high], arr[i+1]
// 	return i + 1
// }

// func quickSort_it(arr []int) {
// 	stack := []int{}
// 	low := 0
// 	high := len(arr) - 1

// 	stack = append(stack, low)
// 	stack = append(stack, high)

// 	for len(stack) > 0 {
// 		high = stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]
// 		low = stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]

// 		pivot := partition(arr, low, high)

// 		if pivot-1 > low {
// 			stack = append(stack, low)
// 			stack = append(stack, pivot-1)
// 		}

// 		if pivot+1 < high {
// 			stack = append(stack, pivot+1)
// 			stack = append(stack, high)
// 		}
// 	}
// }


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
	fmt.Println(nums)

	
	nums = []int{4, 2, 7, 1, 9, 3, 5, 6, 8}

	nums = quickSort(nums)

	fmt.Println(nums)


}
