// Class

// Go does not have a keyword class. But it does not limit us to
// define a class in Go. We can define a class in a way which is
// similar to defining variables; and define methods of a class, by
// providing set of methods on the common type.

// For example, we can define a class by defining a struct and implement
// methods on the struct type. And this struct will be analogous to a class
// concept.

package main

import "fmt"

type student struct {
	name string
	id   int
}

func (s student) PrintDetails() {
	fmt.Printf("\nStudent %d details\n........\n", s.id)
	fmt.Println("Name :", s.name)
	fmt.Println("ID :", s.id)
}

func main() {
	var stud1 = student{name: "clinton", id: 1}
	stud2 := student{name: "Goku", id: 2}
	stud3 := student{name: "Sunmisola", id: 3}

	stud1.PrintDetails()
	stud2.PrintDetails()
	stud3.PrintDetails()
}
