// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

type SyncMap struct {
	// Блокировка на n читателей и 1 писателя
	rwMutex sync.RWMutex
	data    map[string]int
}

// Получение данных из мапы
func (sm *SyncMap) get(k string) (v int, ok bool) {
	// Блокировка мапы на чтение
	sm.rwMutex.RLock()
	// Разблокировка мапы на чтение по завершении функции
	defer sm.rwMutex.RUnlock()
	// Чтение данных из мапы
	v, ok = sm.data[k]
	// Возврат значения и сигнификатора удачного выполнения
	return v, ok
}

// Запись данных в мапу
func (sm *SyncMap) insert(k string, v int) {
	// Блокировка мапы на чтение и запись
	sm.rwMutex.Lock()
	// Разблокировка мапы на чтение и запись по завершении функции
	defer sm.rwMutex.Unlock()
	// Запись, перезапись значения в мапу
	sm.data[k] = v
}

func main() {
	sm := SyncMap{
		data: make(map[string]int),
	}
	sm.insert("donut", 5)
	fmt.Println(sm.get("donut"))

	for i := 0; i < 100; i++ {
		go func(n int) {
			sm.insert("donut", n)
		}(i)
		go func(n int) {
			sm.insert("pancakes", n)
		}(i)
		go func() {
			fmt.Println(sm.get("donut"))
		}()
		go func() {
			fmt.Println(sm.get("pancakes"))
		}()
	}

}
