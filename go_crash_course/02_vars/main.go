package main

import "fmt"

func main() {
	// MAIN Types
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 unintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	//complex64 complex128

	// Var Keyword
	var name = "krupesh"
	var age = 23
	var isCool = true

	// Constant Variables
	const pi = 3.141

	// Shorthand variables
	// use := only for new variables
	surname := "anadkat"
	surname = "Anadkat"

	// Multiple Variable
	firstName, lastName := "Mahesh", "Anadkat"

	fmt.Println(name, surname, age, isCool)
	fmt.Println(firstName, lastName)
	// fmt.Printf("%T\n", name)
}
