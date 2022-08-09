package main

import "fmt"

func main() {
	a := 10
	b := 3

	//logical and "&&"
	if(a != b && a >= b){
		fmt.Println("YES")
	}

	// simple assignment operator "="
	a = b
	fmt.Println(a)
}

