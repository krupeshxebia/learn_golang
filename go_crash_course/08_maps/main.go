package main

import "fmt"

func main() {
	fmt.Println("hello World")
	// // Maps
	// emails := make(map[string]string)

	// // Insert Key - Value
	// emails["krupesh"] = "kmanadkat@gmail.com"
	// emails["mahesh"] = "mmanadkat@gmail.com"
	// emails["mike"] = "mike@gmail.com"

	// Declare and insert key value
	emails := map[string]string{
		"krupesh" : "kmanadkat@gmail.com",
		"mahesh" : "mmanadkat@gmail.com",
		"mike" : "mike@gmail.com",
	}

	// Delete From Map
	delete(emails, "mike")

	fmt.Println(emails)
	fmt.Println(emails["krupesh"])
	fmt.Println(len(emails))
}