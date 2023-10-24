// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
package main

import (
	"fmt"
	"sync"
)

func calculateSquare(num int, wg *sync.WaitGroup, sum *int, m *sync.Mutex) {
	defer wg.Done()
	square := num * num
	m.Lock() //убедился что без мьютекса результат не стабильный видимо когда горутины пытаются записать значения в переменную конкурентно проихсодит конфликт. состоянию гонки race condition
	*sum += square
	m.Unlock()
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	var sum int
	var wg sync.WaitGroup
	var mut sync.Mutex
	for _, num := range nums {
		wg.Add(1)
		go calculateSquare(num, &wg, &sum, &mut)
	}

	wg.Wait()

	fmt.Printf("Sum of squares: %d\n", sum)
}