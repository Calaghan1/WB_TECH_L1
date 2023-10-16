package main

import ("fmt"
		"sync"
)

func calculateSquare(num int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()
	square := num * num
	resultChan <- square
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	resultChan := make(chan int, len(nums))
	var wg sync.WaitGroup

	for _, num := range nums {
		wg.Add(1)
		go calculateSquare(num, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for square := range resultChan {
		fmt.Printf("Square: %d\n", square)
	}
}