// Разработать программу, которая будет последовательно отправлять значения в канал,
//  а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

package main

import (
	"fmt"
	"sync"
	"time"
)
func main() {
	// Создаем канал для обмена данными
	ch := make(chan int)
	var wg sync.WaitGroup
	// Запускаем горутину для чтения из канала
	var t int
	fmt.Println("Сколько секунд работаем?")
	fmt.Scanf("%d\n", &t)
	wg.Add(1)
	to := time.After(time.Duration(t) * time.Second)
	go func() {
		for {
			select {
			case val, ok := <-ch:
			if !ok {
				wg.Done()
				return
			}
				fmt.Printf("Received: %d\n", val)
			}
		}
	}()

	// Запускаем горутину для отправки значений в канал
	wg.Add(1)
	go func() {
		i := 0
		for true {
			select {
				case <-to:
					close(ch)
					wg.Done()
					return
				default:
					ch <- i
					time.Sleep(time.Second) // Ожидаем 1 секунду
					i++
			}

		}
	}()
		wg.Wait()

}