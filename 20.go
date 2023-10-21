// Разработать программу, которая переворачивает слова в строке. 
// Пример: «snow dog sun — sun dog snow».
package main

import (
	"fmt"
	"strings"
)

func reverseWords(input string) string {
	words := strings.Fields(input)
	for i := 0; i < len(words) / 2; i++ {
		words[i], words[i - len(words)] = words[i - len(words)], words[i]
	}
	return strings.Join(words, " ")
}

func Exercise_20() {
	str := "snow dog sun — sun dog snow"
	reversed := reverseWords(str)
	fmt.Printf("Исходная строка: %s\n", str)
	fmt.Printf("Строка с перевернутыми словами: %s\n", reversed)
}