// Реализовать собственную функцию sleep.
package main

import (
	"runtime"
	"time"
)

func Sleep_1(d time.Duration) {
	ticker := time.NewTicker(d)
	<- ticker.C
}

func Exercise_25() {

	runtime.Gosched()
}

