// Реализовать конкурентную запись данных в map.


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
			}(i) //передача i анонимную функцию
		}
	
		// Ожидаем завершения всех горутин
		wg.Wait()
	
		// Выводим данные из map
		fmt.Println(data)
}
func second_7() {
	var m sync.Map
	var wg sync.WaitGroup
	writeToMap := func(key string, value int, wg *sync.WaitGroup) { //ну вот так еще можно функции объявлять) 
		defer wg.Done()
		m.Store(key, value)
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeToMap(fmt.Sprintf("key%d", i), i, &wg)
	}
	wg.Wait()
	m.Range(func(key, value interface{}) bool { //Range вызывает фунцию из своего аргумепнта для каждой пары ключ значение
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true // Возвращаем true, чтобы продолжить итерацию
	})
	
}
func main() {
	first_7()
	second_7()
}