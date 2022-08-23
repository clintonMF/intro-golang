// Go Ranges
// In Go programming, range keyword is used in conjunction with Go For Loop
// to iterate over the items of a collection.

// Range can be used with an array, string, map or channels. During each
// iteration, range can return one or two values based on the type of
// collection you are iterating.

// In the following table, an overview of what range returns during iteration
// is given.

// Range Expression		1st Value	2nd Value
// Array or slice		index		element
// String				index		rune int
// 									(int32 value representing the character value)
// Map					key			value
// Channel				element		none

// Using second value returned by the range is optional. Going forward in this
// article, we will learn how to use golang range with different range
// expressions mentioned in the above table.

package main

import "fmt"

// function a shows the use of range over an array
func a() {
	var even = []int{2, 4, 6}

	// using the first value (index) only
	fmt.Println("using the first value (index) only\n.........")
	for index := range even {
		fmt.Printf("Even[%d] = %d \n", index, even[index])
	}

	fmt.Println("\nusing the first and second value \n.........")
	for index, element := range even {
		fmt.Printf("Even[%d] = %d\n", index, element)
	}
}

// function b shows the use of range over a string

func b() {
	str := "STRING"

	// using the first value only
	fmt.Println("\nusing the fisrt value only")
	for index := range str {
		fmt.Printf("str[%d] = %d\n", index, str[index])
	}

	// using the first and second value
	fmt.Println("\nusing the fisrt value and second")
	for index, element := range str {
		fmt.Printf("str[%d] = %d\n", index, element)
	}
}

// function c shows how to use a range over a map
func c() {
	gamers := map[string]string{"clinton": "PES", "jerry": "FIFA"}

	// using the key only
	fmt.Println("\nUsing the key only\n..................")
	for key := range gamers {
		fmt.Println(key, "plays", gamers[key])
	}

	// using the key and the value
	fmt.Println("\nUsing the key and the value\n..............")
	for key, value := range gamers {
		fmt.Println(key, "plays", value)
	}

}

func main() {
	a()
	b()
	c()
}
