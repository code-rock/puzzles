// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской 
// структуры Human (аналог наследования)

package main

type Human structure {
	wellFed bool
}

// Метод структур Human
// Дополнительным аргументом передается указатель нас структуру
func (h *Human) Eat() {
	fmt.Println("Om nom nom")
	// Изменение свойства структуры Human
	h.wellFed = true
}

type Action structure {
	human Human
}

