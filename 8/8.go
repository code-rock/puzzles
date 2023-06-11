// Дана переменная int64. Разработать программу которая устанавливает
// i-й бит в 1 или 0.

package main

import (
	"flag"
	"fmt"
	"log"
)

func setBit(num *int64, pos int, bit int64) {
	// 1 сдвигаем на необходимое количетсво позиций влево зануляет эту позицию. Сброс бита
	// Бит на который надо заменит сдвигаем на необходимое количество позиций в лево
	// или зануленное значение или значение бита в нужной позиции
	*num = (*num & ^(1 << pos)) | (bit << pos)
}

func main() {
	// Объявление типов переменных
	var num int64
	var pos int
	var bit int64
	// Запись в объявленные переменные соответствующих значений из аргументов команднойстроки
	// Исходное число
	flag.Int64Var(&num, "n", 0, "original number")
	// от 0 до 64
	flag.IntVar(&pos, "p", 0, "bit position")
	// 0 или 1
	flag.Int64Var(&bit, "b", 0, "bit value")
	// Разбора командной строки на определенные флаги
	flag.Parse()
	// Проверка неподходящих значений аргументов
	if pos < 0 || pos > 64 || bit < 0 || bit > 1 {
		log.Fatalln("invalid arguments")
	}
	// Вызов функции установки нужного бита
	setBit(&num, pos, bit)
	// Вывод полученного числа
	fmt.Println(num)
}

// go run 8.go -n=10 -b=1 -p=12