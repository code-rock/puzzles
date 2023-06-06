// Удалить i-ый элемент из слайса.
package main

func remove(arr []string, i int) {
	last := len(arr) - 1
	arr[i] = arr[last]
	arr[last] = ""
	return a[:last]
}

func removeKeepingOrder(arr []string, i int) {
	last := len(arr) - 1
	copy(arr[i:], a[i+1:])
	arr[last] = ""
	return a[:last]
}
