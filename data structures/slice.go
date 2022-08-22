// Golang Slice is an abstraction over Array. While Arrays cannot be
// expanded or shrinked in size, Slices are dynamically sized. Golang
// Slice provides many inbuilt functions, which we will go through
// them in this tutorial.

// To declare a golang slice, use var keyword with variable name
// followed by []T where T denotes the type of elements that will be
// tored in the slice.
// var slicename []T

package main

import "fmt"

// function a declares and initializes a slice
func a() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(numbers)
	fmt.Println("above is function a\n  ")
}

// function b shows how to get element from a slice
func b() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	num1 := numbers[0]
	fmt.Println("the first number:", num1)
	fmt.Println("the second number", numbers[1])
	fmt.Println("above is function b\n  ")
}

// function c deals with length and capacity of a slice
// Length and Capacity of a Slice
// Slice has two properties regarding the number of elements present  in
// slice (length) and the number of elements it can accommodate (capacity).

func c() {
	numbers := make([]int, 5, 10)
	fmt.Println("size of slice", len(numbers))
	fmt.Println("capacity of slice", cap(numbers))
	fmt.Println("above is function c\n  ")

}

// shows append function in slice
func d() {
	numbers := []int{1, 2, 3, 4, 5}

	numbers = append(numbers, 60)
	fmt.Println(numbers)
}

func main() {
	a()
	b()
	c()
	d()
}
