package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tastCases := []struct {
		arr  []int
		want []int
	}{
		{
			arr:  []int{3, 23, 1, 5, 12, 83, 4, 44},
			want: []int{1, 3, 4, 5, 12, 23, 44, 83},
		},
		{
			arr:  []int{},
			want: []int{},
		},
		{
			arr:  []int{8},
			want: []int{8},
		},
		{
			arr:  []int{3, 23, 1, 0, 5, 12},
			want: []int{0, 1, 3, 5, 12, 23},
		},
		{
			arr:  []int{-3, 23, 1, 0, -5, 12},
			want: []int{-5, -3, 0, 1, 12, 23},
		},
	}

	for id, ts := range tastCases {
		if QuickSort(ts.arr); !reflect.DeepEqual(ts.arr, ts.want) {
			t.Errorf("%d QuickSort() = %v, want %v", id+1, ts.arr, ts.want)
		}
	}
}
