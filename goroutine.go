// A goroutine is a lightweight thread managed by the Go runtime.

// To start a goroutine, use keyword go before making the function call as
// shown in the following.

// 			go somefunc(arg1, arg2)

// When the above statement is executed, the statement itself gets executed
// in the current goroutine, but somefunc(arg1, arg2) gets executed in a new
// goroutine.

package main

import "fmt"

func PrintDetails(s string) {
	fmt.Println(s)
}

func main() {
	// using a go routine
	go PrintDetails("Hello!!!")
	// another go routine
	go PrintDetails("Welcome to Golang Goroutines ")

	fmt.Println("End of the main goroutine")
}
