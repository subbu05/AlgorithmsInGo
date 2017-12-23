package main

import "fmt"

func main() {
	n := 5
	result := 0
	for n > 0 {
		result += n
		n--
	}

	fmt.Println(result)

}
