// Реализовать паттерн «адаптер» на любом примере.
package main

import "fmt"

// Adaptee
type MemoryCard struct{}

func (mc *MemoryCard) insert() string {
	return "Карта памяти успешно вставлена!"
}

func (mc *MemoryCard) copyData() string {
	return "Данные скопированы на компьютер!"
}

// Interface
type USB interface {
	connectWithUsbCable()
}

func connectWithUsbCable() {
	fmt.Println("Подключение по кабелю осуществлено!")
}

// Adapter
type CardReader struct {
	*MemoryCard
}

func NewAdapter(mc *MemoryCard) USB {
	return &CardReader{mc}
}

func (cr *CardReader) connectWithUsbCable() {
	fmt.Println(cr.insert(), cr.copyData())
}

func main() {
	cr := CardReader{
		MemoryCard: &MemoryCard{},
	}
	cr.connectWithUsbCable()
}
