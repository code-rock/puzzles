// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

package main

import (
	"reflect"
	"testing"
)

func TestMakeSet(t *testing.T) {
	test := map[string]struct{}{
		"cat":  struct{}{},
		"dog":  struct{}{},
		"tree": struct{}{},
	}
	got := MakeSet([]string{"cat", "cat", "dog", "cat", "tree"})
	if !reflect.DeepEqual(got, test) {
		t.Errorf("MakeSet() = %v, want %v", got, test)
	}
}
