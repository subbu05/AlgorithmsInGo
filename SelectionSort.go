package main

import "fmt"

func main() {

	a := []int{5, 1, 3, 2, 4}

	n := len(a)

	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		temp := a[min]
		a[min] = a[i]
		a[i] = temp
	}

	fmt.Println(a)
}
