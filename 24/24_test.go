package main

import (
	"testing"
)

func TestDistance(t *testing.T) {
	tastCases := []struct {
		p1   Point
		p2   Point
		want float64
	}{
		{p1: Point{x: 1, y: 2}, p2: Point{x: 34, y: 2}, want: 33},
		{p1: Point{x: -1, y: -2}, p2: Point{x: 34, y: 2}, want: 35.22782990761707},
		{p1: Point{x: -1, y: -2}, p2: Point{x: 34, y: 0}, want: 35.05709628591621},
		{p1: Point{x: 0, y: 5}, p2: Point{x: 12, y: 0}, want: 13},
		{p1: Point{x: 0, y: 0}, p2: Point{x: 0, y: 0}, want: 0},
	}
	for _, ts := range tastCases {
		if got := Distance(ts.p1, ts.p2); got != ts.want {
			t.Errorf("distance() = %v, want %v", got, ts.want)
		}
	}
}
