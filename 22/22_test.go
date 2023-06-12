package main

import (
	"math/big"
	"testing"
)

func TestDistance(t *testing.T) {
	tastCases := []struct {
		variables          Calculator
		multiplicationWant *big.Int
		divisionWant       *big.Int
		summarizeWant      *big.Int
		subtractWant       *big.Int
	}{
		{
			variables: Calculator{
				a: big.NewInt(1 << 10),
				b: big.NewInt(1 << 10),
			},
			multiplicationWant: big.NewInt(1 << 20),
			divisionWant:       big.NewInt(1),
			summarizeWant:      big.NewInt(1 << 11),
			subtractWant:       big.NewInt(0),
		},
	}
	for _, ts := range tastCases {
		if got := Multiplication(ts.variables); got.Cmp(ts.multiplicationWant) != 0 {
			t.Errorf("Multiplication() = %v, want %v", got, ts.multiplicationWant)
		}
	}

	for _, ts := range tastCases {
		if got := Division(ts.variables); got.Cmp(ts.divisionWant) != 0 {
			t.Errorf("Division() = %v, want %v", got, ts.divisionWant)
		}
	}

	for _, ts := range tastCases {
		if got := Summarize(ts.variables); got.Cmp(ts.summarizeWant) != 0 {
			t.Errorf("Summarize() = %v, want %v", got, ts.summarizeWant)
		}
	}

	for _, ts := range tastCases {
		if got := Subtract(ts.variables); got.Cmp(ts.subtractWant) != 0 {
			t.Errorf("SummSubtractarize() = %v, want %v", got, ts.subtractWant)
		}
	}
}
