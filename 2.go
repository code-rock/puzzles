package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"time"
)

// Программа, которая конкурентно рассчитает значение
// квадратов чисел взятых из массива (2,4,6,8,10) и выведет их
// квадраты в stdout.

// Создание функции с передоваемым агрументом ns - масив целых чисел
func Square(ns []int) {
	// Создание целочисленного пустого канала
	pows := make(chan int)

	// Итерация по всем элементам масива целых чисел
	for _, value := range ns {
		// Самовызывающаяся горутина, куда аргументом передается целое число
		go func(n int) {
			// Возедение числа в квадрат
			//int(math.Pow(float64(n), 2)) // n * n  invalid composite literal type int
			// блокировка, пока данные не будут получины
			// pow := fmt.Sprintf("%d", math.Pow(float64(n), 2))
			// io.WriteString(os.Stdout, pow)
			pows <- int(math.Pow(float64(n), 2))
		}(value)
	}

	// Вывод степеней в стандартный поток вывода
	// fmt.Printf("%d!", <-pows)
	time.Sleep(3 * time.Second)
	fmt.Printf("%d!", fmt.Sprintf("%v", <-pows))
	io.WriteString(os.Stdout, fmt.Sprintf("%d", <-pows))
}

func main() {
	// Создание масива целых чисел
	ns := []int{2, 3, 8, 12, 54, 1, 0}
	// Вызов функции возведения в квадрат
	Square(ns)
}
