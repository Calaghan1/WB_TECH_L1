package main

import (
	"fmt"
	"sync"
)

func calculateSquare(num int, wg *sync.WaitGroup, sum *int) {
	defer wg.Done()
	square := num * num
	*sum += square
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	var sum int
	var wg sync.WaitGroup

	for _, num := range nums {
		wg.Add(1)
		go calculateSquare(num, &wg, &sum)
	}

	wg.Wait()

	fmt.Printf("Sum of squares: %d\n", sum)
}