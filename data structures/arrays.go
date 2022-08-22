// In Go programming language, Array is a collection of elements of the same datatype.

// Go array is of fixed size. The size mentioned during the declaration or inferred during initialization is the size of the array.

// The array is sequential which means that the elements of the array can be access sequentially.

// Genarally in programming, an Array is used to store similar kind of values or objects where the sequence of elements matters.

// To declare an array in Golang, use the following syntax.
// var array_name [size] datatype

package main

import "fmt"

// function a follow the traditional way of setting array values
// i.e it only declares the array
func a() {
	var primes [5]int
	primes[0] = 1
	primes[1] = 2
	primes[2] = 3
	primes[3] = 5
	primes[4] = 7

	fmt.Println(primes)
}

//function b declares and initialize array values
func b() {
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes)
}

// function c shows how to use array values
func c() {
	primes := [5]int{2, 3, 5, 7, 11}

	fmt.Println("here is the first prime number:", primes[0])
	fmt.Printf("here is the second prime number: %d", primes[1])
}

func main() {
	a()
	b()
	c()
}
