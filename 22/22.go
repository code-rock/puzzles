// Разработать программу, которая перемножает, делит, складывает, вычитает
// две числовых переменных a,b, значение которых > 2^20.
package main

import (
	"log"
	"math/big" // реализует арифметику произвольной точности
)

type Calculator struct {
	a *big.Int
	b *big.Int
}

func Multiplication(c Calculator) *big.Int {
	z := big.NewInt(0)
	return z.Mul(c.a, c.b)
}

func Division(c Calculator) *big.Int {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	z := big.NewInt(0)
	return z.Div(c.a, c.b)
}

func Summarize(c Calculator) *big.Int {
	z := big.NewInt(0)
	return z.Add(c.a, c.b)
}

func Subtract(c Calculator) *big.Int {
	z := big.NewInt(0)
	return z.Sub(c.a, c.b)
}

func main() {}
