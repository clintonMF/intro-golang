// Go Structures
// Go Structure is a datatype that allows you to store different properties
// under a single variable name. It is similar to a class, in object oriented
// programming, that has only properties. It is very similar to structure in
// C programming.

package main

import "fmt"

type student struct {
	name string
	id   int
}

// func a() {
// 	// intialize during declaration
// 	var student1 = student{name: "A", id: 14}
// 	var student2 student
// 	// intialize explicitly
// 	student2 = student{name: "B", id: 13}
// }

func a() {
	student1 := student{name: "clinton", id: 1}
	student2 := student{name: "favour", id: 2}

	fmt.Printf("student1 name is %s he has an id of %d\n", student1.name, student1.id)
	fmt.Printf("student name is %s he has an id of %d", student2.name, student2.id)
}

func b() {
	student1 := student{name: "naruto", id: 3}
	// setting the values
	student1.id = 4
	student1.name = "son goku"
	fmt.Println("\nStudent 1 details\n ........")
	fmt.Println("Name : ", student1.name)
	fmt.Println("id : ", student1.id)
}

// associate method to the student struct type
func (s student) PrintDetails() {
	fmt.Println("\nstudent details\n.............")
	fmt.Println("Name :", s.name)
	fmt.Println("ID :", s.id)
}

func c() {
	student5 := student{name: "Gohan", id: 5}
	student5.PrintDetails()
}

func main() {
	a()
	b()
	c()
}
