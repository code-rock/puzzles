// Поменять местами два числа без создания временной переменной.
package main
func change(a interface{}, b interface{}) {
	return b, a = a, b
}