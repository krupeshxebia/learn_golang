package main

import "fmt"

func main(){
	fmt.Println("Hello World")

	ids := []int{12, 34, 25, 16, 19, 95, 93, 35}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// _ in for
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Taking sum
	sum := 0
	for _, id := range ids {
		 sum += id
	}
	fmt.Printf("Sum is : %d\n", sum)

	// Map with ranges
	emails := map[string]string{
		"krupesh" : "kmanadkat@gmail.com",
		"mahesh" : "mmanadkat@gmail.com",
		"mike" : "mike@gmail.com",
	}
	// Print Key Values
	for k, v := range emails {
		fmt.Printf("%s : %s\n", k, v)
	}
	// Print only Keys
	for k := range emails {
		fmt.Printf("Name : %s\n", k)
	}	
}