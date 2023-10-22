// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

// Например: 
// abcd — true
// abCdefAaf — false
// 	aabcd — false

package main

import (
	"fmt"
	"strings"
)

func areAllCharactersUnique(s string) bool {
	// Преобразовываем строку в нижний регистр
	s = strings.ToLower(s)

	// Создаем map для хранения встреченных символов
	seen := make(map[rune]bool)

	// Перебираем символы в строке
	for _, char := range s {
		// Если символ уже был встречен, значит строка не уникальна
		if seen[char] {
			return false
		}

		// Помечаем символ как встреченный
		seen[char] = true
	}

	// Если дошли до этого места, значит все символы уникальны
	return true
}

func Exercise_26() {
	fmt.Println(areAllCharactersUnique("abcd"))      // true
	fmt.Println(areAllCharactersUnique("abCdefAaf")) // false
	fmt.Println(areAllCharactersUnique("aabcd"))     // false
}





