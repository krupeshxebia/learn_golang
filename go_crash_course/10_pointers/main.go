package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)

	// Use * to dereference value
	fmt.Println(*b)

	// Change value with pointer
	*b = 10
	fmt.Println(a)
}
