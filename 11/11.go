// Реализовать пересечение двух неупорядоченных множеств.
package main

func Intersect(a map[string]struct{}, b map[string]struct{}) map[string]struct{} {
	// Создание мапы для хранения пересечения
	intersection := make(map[string]struct{})
	// Перебор ключей множества
	for k := range a {
		// Если во втором множестве находиться такой же ключ
		if _, ok := b[k]; ok {
			// Добовление ко множеству пересечения
			intersection[k] = struct{}{}
		}
	}
	// Возвращаем множество пересечения
	return intersection
}

func main() {}
