// Реализовать бинарный поиск встроенными методами языка.
package main

// На вход получает предположительно отсортированный массив и искомое значение
func BinarySearch(arr []int, x int) (int, bool) {
	// Объявление левого указателя
	left := 0
	// Объявление правого указателя
	right := len(arr)
	// Пока левый указатель больше правого выполнять следующее
	for left < right {
		// Поиск середины между указателями
		mid := left + (right-left)/2
		// Если середина равна искомому значению
		if arr[mid] == x {
			// Возвращаем индекс зачения и сугнификатор удачного поиска
			return mid, true
			// Если значение по середина менбше искомого
		} else if arr[mid] < x {
			// Смещаем левый указатель на следущее после среднего значение
			// Ибо то уже было провеено и не полдошло
			left = mid + 1
			// Если значение по середина больше искомого
		} else {
			// Смещаем правый указатель на середину
			// В минус не смещаем ибо указатели могут попасть на одно непроверенное значение
			// и по условиню выше выйти из цикла преждевременно
			right = mid
		}
	}

	// Возвращение левого указателя и сигнификатора о неудачном поиске
	return left, false
}

func main() {

}
