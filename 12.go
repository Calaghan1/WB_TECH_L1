// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество

package main

import "fmt"

func Exercise_12(){
	elements := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)
	for _, element := range elements {
        set[element] = true
    }
	fmt.Println(set)
}