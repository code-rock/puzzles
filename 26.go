// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

package main

import (
	"fmt"
	"strings"
)

func AreSymbolsUnique(s string) bool {
	sm := strings.ToLower(s)
	var chs map[string]int

	for i := 0; i < len(sm); i++ {
		char := string(sm[i])
		if chs[char] == 0 {
			chs[char] = 1
		} else {
			return false
		}
	}

	return true
}
