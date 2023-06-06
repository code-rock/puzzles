// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

package main

import (
	"testing"
)

func TestAbs(t *testing.T) {
	tastCases := []struct {
		s    string
		want bool
	}{
		{s: "123", want: true},
		{s: "(2))", want: false},
		{s: "23a$", want: true},
		{s: "2553a$", want: false},
		{s: "54235", want: false},
		{s: "", want: true},
		{s: "abcd", want: true},
		{s: "abCdefAaf", want: false},
		{s: "aabcd", want: false},
	}
	for _, ts := range tastCases {
		if got := AreSymbolsUnique(ts.s); got != ts.want {
			t.Errorf("AreSymbolsUnique() = %v, want %v", got, ts.want)
		}
	}
}
