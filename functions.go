package main

import "fmt"

func multi(a, b int) int {
	out := a * b
	return out
}

func main() {
	a := 3
	b := 2
	c := multi(a, b)

	fmt.Println(c)
}
