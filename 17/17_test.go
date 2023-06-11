package main

import "testing"

type sWanted struct {
	id   int
	flag bool
}

func TestBinarySearch(t *testing.T) {
	tastCases := []struct {
		arr  []int
		n    int
		want sWanted
	}{
		{arr: []int{1, 2, 3, 4, 5, 6}, n: 0, want: sWanted{id: 0, flag: false}},
		{arr: []int{1, 2, 3, 4, 5, 6}, n: 1, want: sWanted{id: 0, flag: true}},
		{arr: []int{1, 2, 3, 4, 5, 6}, n: 4, want: sWanted{id: 3, flag: true}},
		{arr: []int{1, 2, 3, 4, 5, 6}, n: 6, want: sWanted{id: 5, flag: true}},
		{arr: []int{1, 2, 3, 4, 5, 6}, n: 7, want: sWanted{id: 6, flag: false}},
		{arr: []int{}, n: 7, want: sWanted{id: 0, flag: false}},
		{arr: []int{1, 2, 3, 4, 5, 6, 7}, n: 7, want: sWanted{id: 6, flag: true}},
		{arr: []int{1, 2, 3, 4, 5, 6, 7}, n: 2, want: sWanted{id: 1, flag: true}},
		{arr: []int{1, 2, 3, 4, 5, 6, 7}, n: 1, want: sWanted{id: 0, flag: true}},
		{arr: []int{1, 2, 3, 4, 5, 6, 7}, n: 12, want: sWanted{id: 7, flag: false}},
	}

	for _, ts := range tastCases {
		pos, ok := BinarySearch(ts.arr, ts.n)
		if pos != ts.want.id {
			t.Errorf("BinarySearch() = %v, want %v", pos, ts.want.id)
		}
		if ok != ts.want.flag {
			t.Errorf("BinarySearch() = %v, want %v", ok, ts.want.flag)
		}
	}
}
