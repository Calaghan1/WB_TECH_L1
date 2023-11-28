// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. 
// Необходима возможность выбора количества воркеров при старте.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func producer(ch chan int, stopChan chan os.Signal) {
	for true {
		select {
			case <- stopChan:
				close(ch)
				return
			default:
				time.Sleep(time.Second / 5)
				ch <- rand.Intn(10)
		}

	}
}
type Worcker struct {
	Name string
} 
func (w *Worcker)listen(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for true {
		select {
			case data, ok := <- ch:
				if !ok {
					return
				} else {
					fmt.Println(w.Name, " Data: ", data)
				}
		}
	}
}

func main() {
	fmt.Println("Сколько воркеров будет работать?")
	var worker_n int 
	fmt.Scan(&worker_n)
	fmt.Println(worker_n)
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)// при срабатывании сигналов завершения os.Interrupt syscall.SIGTERM отправляет os.Signal в канал os.Interrupt при нажатии Ctrl+C syscall.SIGTERM может быть отправлен другими процессами или сценариями для запроса завершения целевого процесса
	var wg sync.WaitGroup
	var i = 0
	arr := make([]Worcker, 0, 3)
	for i < worker_n {
		i++
		arr = append(arr, Worcker{"Worker "+strconv.Itoa(i)})
	}
	ch := make(chan int, 10)
	i = 0
	for i < len(arr) {
		wg.Add(1)
		go arr[i].listen(ch, &wg)
		i ++
	}
	go producer(ch, stopChan)
	wg.Wait()
}