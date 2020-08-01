package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	gender    string
	age       int
	city      string
}

// Value Receiver
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName
}

// Pointer Receiver
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	person1 := Person{
		firstName: "Krupesh",
		lastName:  "Anadkat",
		gender:    "M",
		age:       23,
		city:      "Gurgaon",
	}

	fmt.Println(person1)

	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1.age)
	fmt.Println(person1.greet())
	person1.hasBirthday()
	fmt.Println(person1.age)
}
