// Поменять местами два числа без создания временной переменной.
package main

func Change(a int, b int) (int, int) {
	b, a = a, b
	return a, b
}

func main() {}
