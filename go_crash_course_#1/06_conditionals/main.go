package main

import "fmt"

func main(){

	x := 5
	y := 6

	// If Else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	}  else {
		fmt.Printf("%d is greater than %d\n", x, y)
	}
 
	// Else If
	color := "orange"

	if color == "red" {
		fmt.Printf("Color is %s\n", color)
	} else if color == "blue" {
		fmt.Printf("Color is %s\n", color)
	} else {
		fmt.Println("Color is neither red or blue")
	}

	// Switch case
	color = "blue"
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is neither red or blue")
	}
}