package main

import (
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	tastCases := []struct {
		variables          Calculator
		multiplicationWant int64
		divisionWant       int64
		summarizeWant      int64
		subtractWant       int64
	}{
		{
			variables: Calculator{
				a: int64(math.Pow(2, 10)),
				b: int64(math.Pow(2, 10)),
			},
			multiplicationWant: int64(math.Pow(2, 20)),
			divisionWant:       1,
			summarizeWant:      int64(math.Pow(2, 11)),
			subtractWant:       0,
		},
	}
	for _, ts := range tastCases {
		if got := Multiplication(ts.variables); got != ts.multiplicationWant {
			t.Errorf("Multiplication() = %v, want %v", got, ts.multiplicationWant)
		}
	}

	for _, ts := range tastCases {
		if got := Division(ts.variables); got != ts.divisionWant {
			t.Errorf("Division() = %v, want %v", got, ts.divisionWant)
		}
	}

	for _, ts := range tastCases {
		if got := Summarize(ts.variables); got != ts.summarizeWant {
			t.Errorf("Summarize() = %v, want %v", got, ts.summarizeWant)
		}
	}

	for _, ts := range tastCases {
		if got := Subtract(ts.variables); got != ts.subtractWant {
			t.Errorf("SummSubtractarize() = %v, want %v", got, ts.subtractWant)
		}
	}
}
