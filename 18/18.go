// Реализовать структуру-счетчик, которая будет инкрементироваться в
// конкурентной среде. По завершению программа должна выводить итоговое
// значение счетчика.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	// Используем uint64 для атомарного счетчика
	value atomic.Uint64
}

func (c *Counter) Increment(n uint64) {
	// Атомарно добавляет дельту к счетчику и возвращает новое значение
	c.value.Add(n)
}

func (c *Counter) Load() uint64 {
	// Атомарно загружает и возвращает значение, хранящееся в счетчике
	return c.value.Load()
}

func main() {
	// Создание группы ожидания всех горутин
	var wg sync.WaitGroup
	// Создание счетчика
	counter := Counter{}
	// ПРоходим в цикле 100 знаений
	for i := 0; i < 100; i++ {
		// Добовление горутины в счетчик группы
		wg.Add(1)
		// Вызываем горутину с текущим числом счетчика
		go func(n int) {
			// Говорим что позавершении функции эта горутина сполнена для группы
			defer wg.Done()
			// Прибоавляем 5 умноженное на id к счетчику
			counter.Increment(uint64(5 * n))
		}(i)
	}

	// Ожидание выполнения всех горутин
	wg.Wait()
	// Вывод конечного значения счетчика
	fmt.Println(counter.Load())
}
