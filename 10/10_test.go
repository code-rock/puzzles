package main

import (
	"reflect"
	"testing"
)

func TestCombineByDegrees(t *testing.T) {
	tastCases := []struct {
		t    []float64
		want map[int64][]float64
	}{
		{
			t: []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5},
			want: map[int64][]float64{
				-20: {-25.4, -27.0, -21.0},
				10:  {13.0, 19.0, 15.5},
				20:  {24.5},
				30:  {32.5},
			},
		},
		{
			t: []float64{-25.4, 0, 13.0, -1.7, 15.5, 0.5, -21.0, 32.5},
			want: map[int64][]float64{
				-20: {-25.4, -21.0},
				0:   {0, -1.7, 0.5},
				10:  {13.0, 15.5},
				30:  {32.5},
			},
		},
		{
			t:    []float64{},
			want: map[int64][]float64{},
		},
		{
			t:    []float64{0},
			want: map[int64][]float64{0: {0}},
		},
	}

	for id, ts := range tastCases {
		got := CombineByDegrees(ts.t)
		if !reflect.DeepEqual(got, ts.want) {
			t.Errorf("%d RemoveKeepingOrder() = %v, want %v", id, got, ts.want)
		}
	}
}
