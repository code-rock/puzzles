// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.
package main

func MakeSet(arr []string) map[string]struct{} {
	// Создание пустого сета
	set := make(map[string]struct{})
	// Итераця по всему масиву строк и получение из них самих строк
	for _, s := range arr {
		// Запись в сет строк
		set[s] = struct{}{}
	}
	// Возврат сета
	return set
}

func main() {}
