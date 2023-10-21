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
var justString string

func createHugeString(size int) string {
    // Создаем большую строку
    hugeString := "..." // Здесь создайте большую строку

    // Возвращаем только первые 100 символов
    return hugeString[:100]
}

func someFunc() {
    v := createHugeString(1 << 10)
    justString = v
}

func Exercise_15() {
    someFunc()
}