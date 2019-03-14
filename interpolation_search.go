package main

import "fmt"

func main() {
	fmt.Println(IPSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
}

func IPSearch(a []int, key int) bool {
	low := 0
	high := len(a) - 1

	for low <= high && key >= a[low] && key <= a[high] {
		pos := low + ((key - a[low]) * (high - low) / (a[high] - a[low]))
		if a[pos] == key {
			return true
		}
		if key < a[pos] {
			high -= 1
		} else {
			low += 1
		}
	}

	return false
}
