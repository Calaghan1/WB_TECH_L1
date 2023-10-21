// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
	"time"
)

type counter struct {
	mu      sync.Mutex
	Count int
}

func foo_18(c *counter, n int, wg *sync.WaitGroup ) {
	i := 0
	wg.Add(1)
	defer wg.Done()
	for i < 10 {
		c.mu.Lock()
		c.Count ++
		c.mu.Unlock()
		fmt.Println("Working ", n)
		time.Sleep(3 * time.Second)
	}

}

func Exercise_18() {
	i := 0
	var wg sync.WaitGroup
	c := counter{}
	for i < 4 {
		go foo_18(&c, i, &wg)
	}
	wg.Wait()
	fmt.Println(c.Count)
}