import "strings"

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func Reverse(s string) {
	words := strings.Fields(s)
	newRow := ""
	for i := len(words) - 1; i >= 0; i-- {
		newRow += words[i]
	}

	return newRow
}