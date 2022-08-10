//pointers are known as special variables
// They are used to save the memory address of another variable

package main

import "fmt"

func main() {
	value := 60
	var a *int
	a = &value

	fmt.Println("value = ", value)
	fmt.Println("address of  value = ", &value)
	fmt.Println("value in a = ", a)
}
