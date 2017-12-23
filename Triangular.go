package main

import "fmt"

func FindTri(n int) int {

	if n == 1 {
		return 1
	}

	return n + FindTri(n-1)

}

func main() {

	n := 5

	fmt.Println(FindTri(n))

}
