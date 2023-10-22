// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

func getType(value interface{}) string {
	return reflect.TypeOf(value).String()
}

func main() {
	var val interface{}

	val = "Hello, World"
	fmt.Printf("Type of variable: %s\n", getType(val))

	val = 42
	fmt.Printf("Type of variable: %s\n", getType(val))

	val = true
	fmt.Printf("Type of variable: %s\n", getType(val))

	ch := make(chan int)
	val = ch
	fmt.Printf("Type of variable: %s\n", getType(val))

	type CustomType struct{}
	val = CustomType{}
	fmt.Printf("Type of variable: %s\n", getType(val))
}