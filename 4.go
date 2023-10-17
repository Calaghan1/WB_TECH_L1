package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func producer(ch chan int) {
	for true {
		time.Sleep(time.Second / 5)
		ch <- rand.Intn(10)
	}
}
type Worcker struct {
	Name string
} 
func (w *Worcker)listen(ch chan int) {
	for data := range ch {
		fmt.Println(w.Name, data)
	}
}

func main() {
	var i = 0
	arr := make([]Worcker, 0, 3)
	for i < 10 {
		i++
		arr = append(arr, Worcker{"Worker"+strconv.Itoa(i)})
	}
	ch := make(chan int, 10)
	i = 0
	for i < len(arr) {
		go arr[i].listen(ch)
		i ++
	}
	go producer(ch)
	time.Sleep(time.Hour)
}