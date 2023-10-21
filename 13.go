// Поменять местами два числа без создания временной переменной.
package main

import "fmt"
func Exercise_13() {
	a := 5
	b := 10

	a = a + b
	b = a - b
	a = a - b

fmt.Println("a:", a)
fmt.Println("b:", b)


a = 5
b = 10

a = a ^ b
b = a ^ b
a = a ^ b

fmt.Println("a:", a)
fmt.Println("b:", b)


a = 5
b = 10

a, b = a+b, a
a = a - b

fmt.Println("a:", a)
fmt.Println("b:", b)

a = 5
b = 10

a, b = b, a

fmt.Println("a:", a)
fmt.Println("b:", b)
}