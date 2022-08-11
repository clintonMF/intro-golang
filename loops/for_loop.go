package main

import "fmt"

// there are various types of for loops in golang
//
func main() {
	//for loops with condition
	a := 0
	for a < 5 {
		fmt.Println("a")
		a = a + 1
	}

	//for loops with initialisation, condition and update
	for i := 0; i < 5; i++ {
		fmt.Println("i")
	}

	// for loops over range
	b := [6]int{1, 3, 4, 5, 6, 7}
	for j, x := range b {
		fmt.Printf("b[%d] = %d\n", j, x)
	}

}
