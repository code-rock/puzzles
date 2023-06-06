// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.
package main

import "strings"

func Reverse(s string) {
	words := strings.Fields(s)
	newRow := ""
	for i := len(words) - 1; i >= 0; i-- {
		newRow += words[i]
	}

	return newRow
}
