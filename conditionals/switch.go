package main

import "fmt"

// switch case with expression syntax
// switch expression {
// case value1:
// 	statement(s)
// case value2:
// 	statement(s)
// default:
// 	statement(s)
// }

func main() {
	var today int = 2

	switch today {
	case 1:
		fmt.Println("today is Mondday")
	case 2:
		fmt.Println("today is Tuesday")
	case 3:
		fmt.Println("today is Wednesday")
	case 4:
		fmt.Println("today is Thursday")
	case 5:
		fmt.Println("today is Friday")
	case 6:
		fmt.Println("today is Saturday")
	case 7:
		fmt.Println("today is sunday")
	default:
		fmt.Println("the value for today is invalid")
	}
}

// switch case without expression syntax
// switch {
// case expression1==value1:
// 	statement(s)
// case expression2==value2:
// 	statement(s)
// default:
// 	statement(s)    
// }
