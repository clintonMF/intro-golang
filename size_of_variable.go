// In this file we use the unsafe package to check the size of variables

package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var a int

	// uint -> unsigned integer => it can only take positive integers
	var aUint8 uint8
	// var aUint16 uint16
	// var aUint32 uint32
	// var aUint64 uint64

	// int -> regular integers =. can take positive and negative integers
	var aInt8 int8
	// var aInt16 int16
	// var aInt32 int32
	// var aInt64 int64

	var c float32 // used to store decimal numbers
	var d bool    //True or false values. default value is false
	var e string

	aa := "altschool"

	fmt.Printf("a = %v [%T, %d bits]\n", a, a, unsafe.Sizeof(a))
	fmt.Printf("aUint8 = %v [%T, %d bits]\n", aUint8, aUint8, unsafe.Sizeof(aUint8))
	fmt.Printf("aInt8 = %v [%T, %d bits]\n", aInt8, aInt8, unsafe.Sizeof(aInt8))
	fmt.Printf("c = %v [%T, %d bits]\n", c, c, unsafe.Sizeof(c))
	fmt.Printf("d = %v [%T, %d bits]\n", d, d, unsafe.Sizeof(d))
	fmt.Printf("e = %v [%T, %d bits]\n", e, e, unsafe.Sizeof(e))
	fmt.Printf("aa = %v [%T, %d bits]\n", aa, aa, unsafe.Sizeof(aa))
}
