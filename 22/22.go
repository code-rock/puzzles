// Разработать программу, которая перемножает, делит, складывает, вычитает
// две числовых переменных a,b, значение которых > 2^20.
package main

type Calculator struct {
	a int64
	b int64
}

func Multiplication(c Calculator) int64 {
	return c.a * c.b
}

func Division(c Calculator) int64 {
	return c.a / c.b
}

func Summarize(c Calculator) int64 {
	return c.a + c.b
}

func Subtract(c Calculator) int64 {
	return c.a - c.b
}

