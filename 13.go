// Поменять местами два числа без создания временной переменной.
package main
import "fmt"




func main() {
	a := 5
	b := 10

	a = a + b // Первый способ чисто математический 3 + 2 = 6   6 - 2 = 3 	6 - 3 = 2 работает с целыми числами
	b = a - b
	a = a - b

fmt.Println("a:", a)
fmt.Println("b:", b)


a = 5
b = 10
// "XOR swap
a = a ^ b // в а лежит теперь a ^ b
b = a ^ b// в b теперь лежит (a ^ b) ^ b что по сути есть a
a = a ^ b// в a теперь лежит (a ^ b) ^ (a ^ b) ^ b что есть b

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