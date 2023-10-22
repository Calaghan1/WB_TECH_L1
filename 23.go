// Удалить i-ый элемент из слайса.

package main

import "fmt"

func main() {
    // Исходный слайс
    numbers := []int{1, 2, 3, 4, 5}

    // Индекс элемента, который вы хотите удалить
    indexToRemove := 2 // Удалим элемент с индексом 2 (третий элемент)

    // Удаление элемента из слайса
    numbers = append(numbers[:indexToRemove], numbers[indexToRemove+1:]...)

    // Вывод результата
    fmt.Println(numbers) // Выводит: [1 2 4 5]
}