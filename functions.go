package main

import "fmt"

func multiply(a, b int) int {
	out := a * b
	return out
}

//the function below returns multiple values
func add_multi(a, b int) (int, int) {
	add := a + b
	multi := a * b
	return add, multi
}

func main() {
	a := 3
	b := 2
	c := multiply(a, b)

	fmt.Println(c)

	var d, e int = add_multi(a, b)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Printf("addition = %d \nmultiplication = %d", d, e)
}
