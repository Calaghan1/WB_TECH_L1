// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func first() {
	ch := make(chan bool)
	go func() {
		for true {
		fmt.Println("Gorutine working")
		time.Sleep(3 * time.Second)
		select {
			case <-ch: // Блок кода, выполняется, если можно прочитать из ch
				break
		}
		fmt.Println("Gorutine shot down")
	}
	}()
	time.Sleep(10 * time.Second)
	ch <- true
	time.Sleep(3 * time.Second)
}
func second() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for true {
		fmt.Println("Gorutine working")
		time.Sleep(3 * time.Second)
		select {
			case <-ctx.Done(): // Блок кода, выполняется, если можно прочитать из ch
				break
		}
		fmt.Println("Gorutine shot down")
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
				break 
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
	go func() {
		for true {
			fmt.Println("Gorutine working")
			select {
				case <- sigCh: // Блок кода, выполняется, если можно прочитать из ch
					break
			}
			time.Sleep(1 * time.Second)
			fmt.Println("Gorutine shot down")
		}
	}()
	time.Sleep(10 * time.Second) 
}


func Exercise_6() {

}