// Реализовать пересечение двух неупорядоченных множеств.
package main

import (
	"fmt"
	"math/rand"
)

func Exercise_11() {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)
	i := 0
	for i < 10 {
		i ++
		set1[rand.Intn(10)] = true
		set2[rand.Intn(10)] = true
	}

	set3 := make(map[int]bool)

	for item := range set1 {
		if set2[item] {
			set3[item] = true
		}
	}
	for item := range set3 {
		fmt.Println(item)
	}
}