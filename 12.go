// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество

package main

import "fmt"


func contains(s string, c rune) bool {
	rune_s := []rune(s)
	for _, v := range rune_s {
		if v == c {
			return true
		}
	}
	return false
}


func main(){
	elements := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)
	for _, element := range elements {
			set[element] = true
    }
	fmt.Println(set)

	for k := range set {
		delete(set, k)
	}
	for _, element := range elements {
		if contains(element, 'c'){
			set[element] = true
		}	 
    }
	fmt.Println(set)
}