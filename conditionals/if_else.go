package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	if a > b {
		fmt.Println("a is greater than b")
	} else if a == b {
		fmt.Println("a is equal to b")
	} else {
		fmt.Println("a is less than b")
	}
}
