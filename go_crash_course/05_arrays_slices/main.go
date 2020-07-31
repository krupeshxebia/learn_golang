package main

import "fmt"

func main() {
	// Arrays
	var fruits [2]string
	vegetables := [2]string{"Potato", "Lemon"}

	// Assign Values
	fruits[0] = "Apple"
	fruits[1] = "Watermelon"

	// Print Elements
	fmt.Println(fruits)
	fmt.Println(fruits[1])
	fmt.Println(vegetables)


	// Slices
	vegetableSlice := []string{"Potato", "Lemon", "Cabbage"}

	// 
	fmt.Println(vegetableSlice)
	fmt.Println(vegetableSlice[1:2])
	fmt.Println(len(vegetableSlice))
}
 