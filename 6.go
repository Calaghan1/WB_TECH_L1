// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func first() { //Послать в специальный канал что-то чтобы завершить горутину
	ch := make(chan bool)
	go func() {
		for true {
		select {
			case <- ch:
				fmt.Println("Gorutine shot down") // Блок кода, выполняется, если можно прочитать из ch
				return
			default:
				fmt.Println("Gorutine working")
				time.Sleep(3 * time.Second)
		}
		
	}
	}()
	time.Sleep(10 * time.Second)
	ch <- true
	time.Sleep(3 * time.Second)
}
func second() { //Исползовать контекст
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for true {
		select {
			case <-ctx.Done():
				fmt.Println("Gorutine shot down") // Блок кода, выполняется, если можно прочитать из ch
				return
		default:
			fmt.Println("Gorutine working")
			time.Sleep(3 * time.Second)
			
		
		}
		
	}
	}()
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)

}
func third() {
	go func() {
		i := 0
		for true {
			fmt.Println("Gorutine working")
			if i > 5 {
				return 
			}
			i++
			time.Sleep(1 * time.Second)
			fmt.Println("Gorutine shot down")
		}
	}()
	time.Sleep(10 * time.Second) 
}

func fourth() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		defer wg.Done()
		for true {
			select {
				case <- sigCh: // Блок кода, выполняется, если можно прочитать из ch
				fmt.Println("Gorutine shot down")
					return
			default:
				fmt.Println("Gorutine working")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	wg.Wait()
}

func fifth(t time.Duration) {
	ticker := time.NewTicker(t)
	for {
		select {
		case <-ticker.C:
			fmt.Println("Gorutine shot down")
			return
		default:
			fmt.Println("Gorutine working")
			time.Sleep(1 * time.Second)
		}
		
	}
}
func main() {
	// first()
	second()
	// third()
	// fourth()
	// fifth(3 * time.Second)
}