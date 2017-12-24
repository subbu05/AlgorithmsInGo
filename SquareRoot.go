package main

import "fmt"

func main() {

	x:=49

	r:=x
	for (r*r) > x {
		r=(r+x/r)/2
	}
	fmt.Println(int(r))
}
