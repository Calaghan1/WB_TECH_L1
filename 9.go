// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func Exercise_9() {
	ch := make(chan int)
	ch2 := make(chan int)
	var a [10]int = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	go func() {
		for _, i := range a {
			ch <- i
		}
		close(ch)
	}
	go func () {
		var buff int
		for i := range ch {
			i *= i
			ch2 <- i
		}
		close(ch2)
	}

	for i := range ch2 {
		fmt.Println(i)
	}
}