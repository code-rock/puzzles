// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

package main

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	tastCases := []struct {
		bunchA map[string]struct{}
		bunchB map[string]struct{}
		want   map[string]struct{}
	}{
		{
			bunchA: map[string]struct{}{"cat": struct{}{}, "dog": struct{}{}},
			bunchB: map[string]struct{}{"cat": struct{}{}, "dog": struct{}{}},
			want:   map[string]struct{}{"cat": struct{}{}, "dog": struct{}{}},
		},
		{
			bunchA: map[string]struct{}{},
			bunchB: map[string]struct{}{"cat": struct{}{}, "dog": struct{}{}},
			want:   map[string]struct{}{},
		},
		{
			bunchA: map[string]struct{}{"cat": struct{}{}},
			bunchB: map[string]struct{}{"cat": struct{}{}, "dog": struct{}{}},
			want:   map[string]struct{}{"cat": struct{}{}},
		},
		{
			bunchA: map[string]struct{}{"cat": struct{}{}, "dog": struct{}{}},
			bunchB: map[string]struct{}{"cat": struct{}{}},
			want:   map[string]struct{}{"cat": struct{}{}},
		},
		{
			bunchA: map[string]struct{}{"cat": struct{}{}, "dog": struct{}{}},
			bunchB: map[string]struct{}{"dog": struct{}{}},
			want:   map[string]struct{}{"dog": struct{}{}},
		},
		{
			bunchA: map[string]struct{}{"cat": struct{}{}, "parrot": struct{}{}, "dog": struct{}{}},
			bunchB: map[string]struct{}{"parrot": struct{}{}, "dog": struct{}{}},
			want:   map[string]struct{}{"parrot": struct{}{}, "dog": struct{}{}},
		},
		{
			bunchA: map[string]struct{}{"cat": struct{}{}, "parrot": struct{}{}, "dog": struct{}{}},
			bunchB: map[string]struct{}{"parrot": struct{}{}, "dog": struct{}{}},
			want:   map[string]struct{}{"dog": struct{}{}, "parrot": struct{}{}},
		},
	}

	for id, ts := range tastCases {
		intersection := Intersect(ts.bunchA, ts.bunchB)
		if !reflect.DeepEqual(intersection, ts.want) {
			t.Errorf("%d RemoveKeepingOrder() = %v, want %v", id, intersection, ts.want)
		}
	}
}
