// Разработать конвейер чисел. Даны два канала: в первый пишутся
// числа (x) из массива, во второй — результат операции x*2, после
// чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
	"io"
	"os"
)

// Вывод данных из канала в стандартный поток вывода
func runPrinter(input <-chan int) {
	// Перебор элементов записаные ранее в канал
	for num := range input {
		// Записываем значения в стандартный поток вывода
		io.WriteString(os.Stdout, fmt.Sprintf("%d ", num))
	}
}

// Функция умножения чисел из канала на 2 и запись в другой канал
func runMultiplier(input <-chan int, output chan<- int) {
	// Канал закроется после цикла, ну или ошибки какой, панике
	defer close(output)
	// Перебор элементов записаные ранее в канал
	for num := range input {
		// Умножение значения на два и запись в канал
		output <- num * 2
	}
}

// Функция записи чисел в канал
func runProducer(arr []int, output chan<- int) {
	// Канал закроется после цикла, ну или ошибки какой, панике
	defer close(output)
	// Перебор элементов массива и запись в конал
	for _, num := range arr {
		output <- num
	}
}

func main() {
	// Создание главного кнала
	originalCh := make(chan int)
	// Запуск горутины пишушей данные в канал
	go runProducer([]int{1, 2, 3, 4, 5}, originalCh)
	// Создание канала умножения чисел на 2
	multipliedCh := make(chan int)
	// Данные из закрытого канала и новый канал передаются аргументами в функцию умножения
	go runMultiplier(originalCh, multipliedCh)
	// Данные из закрытого канала передаются в функцию печати
	runPrinter(multipliedCh)
}
