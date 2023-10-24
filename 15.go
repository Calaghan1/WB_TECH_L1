// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

package main

import (
	"fmt"
	"strings"
)
var justString string

func createHugeString(size int) string { // не очень понятно зачем создавать большую строку чтобы потом использовать только 100 первых символов  тем более что v у нас переменная локальная и вся работа точно окажется не нужна
    // Создаем большую строку
    var builder strings.Builder
    i := 0
    for i < size {
        builder.WriteByte(byte(i % 120))
        i ++
    }
    // Здесь создайте большую строку
    hugeString := builder.String()
    // Возвращаем только первые 100 символов
    return hugeString
}

func someFunc() {
    v := createHugeString(100) //раз нужно 100 то можно передать 100
    justString = v

    // justString = createHugeString(100) //ну или так
   

}

func main() {
    someFunc()
    fmt.Println(justString)
}