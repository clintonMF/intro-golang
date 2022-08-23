// Go Map
// In Go, Map is a collection of unordered key:value pairs. While key has
// to be distinct in a Map, value can occur in duplicates.

// You can access the key:value pair of a Map using key. Key in a Map acts
// like an index in an Array.

// To declare a Map, we may use map keyword along with the datatypes of key and value.

// var map_name map[key_data_type]value_data_type
// Please not that the square brackets are mandatory around the key data type.

package main

import "fmt"

// function a shows how to create a map
func a() {
	var colorMap = map[string]string{"white": "FFFFFF", "black": "#OOOOOO"}

	fmt.Println(colorMap)
}

// function b shows another way to create a map
func b() {
	var colorMap = make(map[string]string)

	colorMap["white"] = "#FFFFFF"
	colorMap["black"] = "#OOOOOO"
	fmt.Println(colorMap)
}

// function c shows how to access the key: value pairs of a map
func c() {
	var colorMap = map[string]string{"white": "FFFFFF", "black": "#OOOOOO"}

	fmt.Println(colorMap["white"])
	fmt.Println(colorMap["black"])

}

// function d shows us how to iterate over a map
func d() {
	var colorMap = map[string]string{"white": "FFFFFF", "black": "#OOOOOO"}

	for key, value := range colorMap {
		fmt.Println("Hex value of key", key, "is", value)
	}
}

// function e shows how to update and delete key value of a map
func e() {
	var colorMap = map[string]string{"white": "FFFFFF", "black": "#OOOOOO"}

	colorMap["white"] = "#EEEEEEE"
	delete(colorMap, "black")

	fmt.Println(colorMap)
	// the value of white would change from "#FFFFFF" to "#EEEEEEE"
	// while the key value of black would not exist

}

func main() {
	a()
	b()
	c()
	d()
	e()
}
