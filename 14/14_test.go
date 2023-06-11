// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

package main

import (
	"testing"
)

func TestDetermineType(t *testing.T) {
	tastCases := []struct {
		arg  interface{}
		want string
	}{
		{arg: "", want: "string"},
		{arg: "3f3", want: "string"},
		{arg: make(chan int), want: "unknown"},
		{arg: false, want: "bool"},
		{arg: true, want: "bool"},
		{arg: 0, want: "int"},
		{arg: 1, want: "int"},
		{arg: 12, want: "int"},
		{arg: 13e243, want: "unknown"},
		{arg: 133.00, want: "unknown"},
		{arg: 133.30, want: "unknown"},
	}

	for _, ts := range tastCases {
		if got := DetermineType(ts.arg); got != ts.want {
			t.Errorf("DetermineType() = %v, want %v", got, ts.want)
		}
	}
}

func TestDetermineGenericType(t *testing.T) {
	tastCases := []struct {
		arg  interface{}
		want string
	}{
		{arg: "", want: "string"},
		{arg: "3f3", want: "string"},
		{arg: make(chan int), want: "channel"},
		{arg: false, want: "bool"},
		{arg: true, want: "bool"},
		{arg: 0, want: "int"},
		{arg: 1, want: "int"},
		{arg: 12, want: "int"},
		{arg: 13e243, want: "float"},
		{arg: 133.00, want: "float"},
		{arg: 133.30, want: "float"},
	}

	for _, ts := range tastCases {
		if got := DetermineGenericType(ts.arg); got != ts.want {
			t.Errorf("DetermineGenericType() = %v, want %v", got, ts.want)
		}
	}
}
