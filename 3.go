package main

import (
	"fmt"
	"sync"
)

func calculateSquare(num int, wg *sync.WaitGroup, sum *int, m *sync.Mutex) {
	defer wg.Done()
	square := num * num
	m.Lock()
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