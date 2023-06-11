package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Res struct {
	arr []string
	err error
}

func TestRemove(t *testing.T) {
	errorText := fmt.Errorf("Такого элемента нет в массиве")
	tastCases := []struct {
		arr  []string
		i    int
		want Res
	}{
		{arr: []string{}, i: 0, want: Res{arr: nil, err: errorText}},
		{arr: []string{"5", "6", "8"}, i: -4, want: Res{arr: nil, err: errorText}},
		{arr: []string{"5"}, i: 0, want: Res{arr: []string{}, err: nil}},
		{arr: []string{"5", "6", "8"}, i: 1, want: Res{arr: []string{"5", "8"}, err: nil}},
		{arr: []string{"5", "6", "8", "12"}, i: 1, want: Res{arr: []string{"5", "12", "8"}, err: nil}},
		{arr: []string{"5", "6", "8", "12"}, i: 0, want: Res{arr: []string{"12", "6", "8"}, err: nil}},
		{arr: []string{"5", "6", "8", "12"}, i: 3, want: Res{arr: []string{"5", "6", "8"}, err: nil}},
	}
	for _, ts := range tastCases {
		arr, err := Remove(ts.arr, ts.i)
		if err != nil {
			assert.Equal(t, err, ts.want.err, fmt.Sprintf("Remove() = %v, want %v", err, ts.want.err))
		}
		if arr != nil {
			if !reflect.DeepEqual(arr, ts.want.arr) {
				t.Errorf("Remove() = %v, want %v", arr, ts.want.arr)
			}
		}
	}
}

func TestRemoveKeepingOrder(t *testing.T) {
	errorText := fmt.Errorf("Такого элемента нет в массиве")
	tastCases := []struct {
		arr  []string
		i    int
		want Res
	}{
		{arr: []string{}, i: 0, want: Res{arr: nil, err: errorText}},
		{arr: []string{"5", "6", "8"}, i: -4, want: Res{arr: nil, err: errorText}},
		{arr: []string{"5"}, i: 0, want: Res{arr: []string{}, err: nil}},
		{arr: []string{"5", "6", "8"}, i: 1, want: Res{arr: []string{"5", "8"}, err: nil}},
		{arr: []string{"5", "6", "8", "12"}, i: 1, want: Res{arr: []string{"5", "8", "12"}, err: nil}},
		{arr: []string{"5", "6", "8", "12"}, i: 0, want: Res{arr: []string{"6", "8", "12"}, err: nil}},
		{arr: []string{"5", "6", "8", "12"}, i: 3, want: Res{arr: []string{"5", "6", "8"}, err: nil}},
	}

	for _, ts := range tastCases {
		arr, err := RemoveKeepingOrder(ts.arr, ts.i)
		if err != nil {
			assert.Equal(t, err, ts.want.err, fmt.Sprintf("RemoveKeepingOrder() = %v, want %v", err, ts.want.err))
		}
		if arr != nil {
			if !reflect.DeepEqual(arr, ts.want.arr) {
				t.Errorf("RemoveKeepingOrder() = %v, want %v", arr, ts.want.arr)
			}
		}
	}

}
