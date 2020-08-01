package main

import "fmt"

func main() {
	//  For Loop
	for i := 0; i < 10; i++ {	
		fmt.Printf("Number is %d\n", i)
	}

	// FizzBuzz
	for i := 1; i<= 100; i++ {
		if i % 15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
