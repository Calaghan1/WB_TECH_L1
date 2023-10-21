package main

import (
	"fmt"
	"sync"
)

func first_7() {
		// Создаем map, в которой будем хранить данные
		data := make(map[string]int)

		// Создаем мьютекс для синхронизации доступа к map
		var mutex sync.Mutex
	
		// Количество горутин, которые будут записывать данные в map
		numGoroutines := 10
	
		// Создаем WaitGroup для ожидания завершения всех горутин
		var wg sync.WaitGroup
		wg.Add(numGoroutines)
	
		for i := 0; i < numGoroutines; i++ {
			go func(i int) {
				defer wg.Done()
				key := fmt.Sprintf("key%d", i)
				value := i * 10
	
				// Захватываем мьютекс перед записью в map
				mutex.Lock()
				data[key] = value
				mutex.Unlock()
			}(i)
		}
	
		// Ожидаем завершения всех горутин
		wg.Wait()
	
		// Выводим данные из map
		fmt.Println(data)
}
func second_7() {
	var map sync.Map
	writeToMap := func(key string, value int) {
		map.Store(key, value)
	}
	for i := 0; i < 10; i++ {
		go writeToMap(fmt.Sprintf("key%d", i), i)
	}
	fmt.Scanln()
}
func Exercise_7() {

}