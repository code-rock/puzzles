// Программа, которая конкурентно рассчитает значение
// квадратов чисел взятых из массива (2,4,6,8,10) и выведет их
// квадраты в stdout.

package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

func Square(arr []int) {
	// Объявление группы горутин, которые должны выполняться вместе
	var wg sync.WaitGroup

	// Перебор элементов массива чисел
	for i := range arr {
		// Запись адреса iого элемента массива в переменную
		num := &arr[i]
		// Добовление к внутренниму счетчику группы одну активную горутину
		wg.Add(1)
		// Создание горутины
		go func() {
			// Добовление в стек вызовов функии функции сигнализирующей группе, что горутина завершена
			// Стек вызова будет запущен при выходе из функции
			defer wg.Done()
			// Перезапись значения в исходном массиве на его квадрат
			*num *= *num
		}()
	}
	// Ожидание выполнения всех горутин группы
	wg.Wait()
}

func main() {
	// Создание масива целых чисел
	ns := []int{2, 3, 8, 12, 54, 1, 0}
	// Вызов функции возведения в квадрат
	Square(ns)
	// Вывод степеней в стандартный поток вывода
	io.WriteString(os.Stdout, fmt.Sprintf("%v", ns))
}
