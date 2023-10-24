// Реализовать собственную функцию sleep.
package main

import (
	"time"
	"os/exec"
	"fmt"
)

func Sleep_1(d time.Duration) {
	ticker := time.NewTicker(d)
	<- ticker.C
}

func Sleep_2(t int) {
	com := exec.Command("sleep", fmt.Sprint(t))
	com.Run()
}
func main() {
	Sleep_1(3 * time.Second)
	fmt.Println("I AM HERE ")
	Sleep_2(13)
}

