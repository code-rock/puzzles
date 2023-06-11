package main

import "testing"

func TestReverse(t *testing.T) {
	tastCases := []struct {
		s    string
		want string
	}{
		{s: "snow dog sun", want: "sun dog snow"},
		{s: "   fd ss", want: "ss fd"},
		{s: "   fd   ss", want: "ss fd"},
		{s: "   ", want: ""},
	}

	for _, ts := range tastCases {
		if got := Reverse(ts.s); got != ts.want {
			t.Errorf("Reverse() = %v, want %v", got, ts.want)
		}
	}
}
