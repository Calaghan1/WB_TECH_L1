// Удалить i-ый элемент из слайса.

package main

import "fmt"

func main() {
    // Исходный слайс
    numbers := []int{1, 2, 3, 4, 5}

    // Индекс элемента, который вы хотите удалить
    indexToRemove := 2 // Удалим элемент с индексом 2 (третий элемент)

    // Удаление элемента из слайса
    numbers = append(numbers[:indexToRemove], numbers[indexToRemove+1:]...) // Если важен порядок то апендим элемант справа от удаленного элемента к левой части

    fmt.Println(numbers) // Выводит: [1 2 4 5]
    //Если порядок не важен то можно просо поменять местами последний элемент и тот котороый нужно удалить
    numbers = []int{1, 2, 3, 4, 5}

    numbers[indexToRemove] = numbers[len(numbers) - 1]
    numbers = numbers[:len(numbers) - 1]
    fmt.Println(numbers)
    // Вывод результата
    
}