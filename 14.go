// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

func getType(value interface{}) string {
	switch reflect.TypeOf(value).Kind() {
	case reflect.String:
		return "string"
	case reflect.Int:
		return "int"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "channel"
	default:
		return "unknown"
	}
}

func Exercise_14() {
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

	// Adding a custom type
	type CustomType struct{}
	val = CustomType{}
	fmt.Printf("Type of variable: %s\n", getType(val))
}