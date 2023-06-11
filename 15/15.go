// К каким негативным последствиям может привести данный фрагмент кода,
// и как это исправить? Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(n int) string {
	return fmt.Sprint(n)
}
func someFunc() {
	// В функцию передается 2 в 10 степени из которой делается большая строка
	v := createHugeString(1 << 10)
	// Содание копии необходимой подстроки, чтобы начальная строка могла очиститься
	justString = strings.Clone(v[:100])
	// Cтроки в go иммутабельны
	// В начальном примере будет хранится изначальная строка с пометкой от какого
	// до какого элемента пренадлежит слайсу
	// justString = v[:100]
}

func main() {
	someFunc()
}
