package main

import "testing"

func TestReverse(t *testing.T) {
	tastCases := []struct {
		s    string
		want string
	}{
		{s: "главрыба", want: "абырвалг"},
		{s: "", want: ""},
		{s: "s", want: "s"},
		{s: "dsd", want: "dsd"},
	}
	for _, ts := range tastCases {
		if got := Reverse(ts.s); got != ts.want {
			t.Errorf("Reverse() = %v, want %v", got, ts.want)
		}
	}
}
