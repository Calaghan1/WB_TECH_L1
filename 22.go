// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
    "fmt"
    "math"
    "math/big"
)

func first() {
    fmt.Println(math.MaxInt64)
    a := int64(2<<20) + 10000000 // Присваиваем значение, большее чем 2^20
    b := int64(2<<20) + 500000  // Присваиваем ещё одно значение, большее чем 2^20
    fmt.Println(2<<20)
    fmt.Println(a)
    // Выполняем арифметические операции
    sum := a + b
    difference := a - b
    product := a * b
    quotient := a / b

    // Выводим результаты
    fmt.Printf("Сумма: %d\n", sum)
    fmt.Printf("Разница: %d\n", difference)
    fmt.Printf("Произведение: %d\n", product)
    fmt.Printf("Частное: %d\n", quotient)
}

func second() {
    a := big.NewInt(2 << 30)
    b := big.NewInt(2 << 60)
    
    sum := new(big.Int)
    sum.Add(a, b)

    difference := new(big.Int)
    difference.Sub(b, a)

    product := new(big.Int)
    product.Mul(b, a)
    quotient := new(big.Int)
    quotient.Div(b, a)
    fmt.Println(sum)
    fmt.Println(difference)
    fmt.Println(product)
    fmt.Println(quotient)
}
func main() {
    second()
}