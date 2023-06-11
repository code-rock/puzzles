package main

import (
	"testing"
)

func TestChange(t *testing.T) {
	tastCases := []struct {
		a     int
		b     int
		wantA int
		wantB int
	}{
		{a: 12, b: 5, wantA: 5, wantB: 12},
		{a: 0, b: 5, wantA: 5, wantB: 0},
		{a: 1, b: 1, wantA: 1, wantB: 1},
	}
	for _, ts := range tastCases {
		gotA, gotB := Change(ts.a, ts.b)
		if gotA != ts.wantA {
			t.Errorf("Change() = %v, want %v", gotA, ts.wantA)
		}
		if gotB != ts.wantB {
			t.Errorf("Change() = %v, want %v", gotB, ts.wantB)
		}
	}
}
