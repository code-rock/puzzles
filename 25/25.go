// Реализовать собственную функцию sleep.
package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) time.Time {
	// Возрат значения текущего времени из канала таймера,
	// Произайдет по истечению передоваемого времени
	return <-time.After(duration)
}

func main() {
	// Вызов функции ожидающей 8с
	fmt.Println(Sleep(8 * time.Second))

	// Печать по проществию времени
	fmt.Println("Sleep Over.....")
}
