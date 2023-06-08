// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской
// структуры Human (аналог наследования)

package main

import "fmt"

// Cтруктура Human с полем сигнализирующим сыт ли он
type Human struct {
	wellFed bool
}

// Метод для структур Human
// Дополнительным аргументом передается указатель на структуру, что дает
// возможно редактировать оригенальную структуру
func (h *Human) Eat() {
	fmt.Println("Om nom nom")
	// Изменение свойства структуры Human
	h.wellFed = true
}

// Создание структуры Action
type Action struct {
	// Наследование полей и методов структуры Human
	Human
	// Дополнительные поля с координатами местоположения
	x float64
	y float64
}

func main() {
	Gena := Human{}
	fmt.Println("Гена покушал", Gena.wellFed)
	Gena.Eat()
	fmt.Println("Гена покушал", Gena.wellFed)

	Oleg := Action{}
	fmt.Println("Олежа покушал", Oleg.wellFed)
	Oleg.Eat()
	fmt.Println("Олежа покушал", Oleg.wellFed)
	fmt.Println("Он сейчас в точке", Oleg.x, Oleg.y)
	Oleg.x = 12
	Oleg.y = 5
	fmt.Println("Он сейчас в точке", Oleg.x, Oleg.y)
}
