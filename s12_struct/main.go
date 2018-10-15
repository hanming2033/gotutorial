package main

import (
	"fmt"
)

func main() {
	// * Struct
	// type person is a struct
	type person struct {
		first string
		last  string
		age   int
	}

	type secretAgent struct {
		// ! if person person(explicit) then fields of person can only be access using secretAgent.person.xxx
		// ! if person(anonymous / embeded) then fields of person can be access using secretAgent.person.xxx or secretAgent.xxx
		person // embed person type into this type, no need to give type details
		ltk    bool
	}
	// create a value of type person using composite literal and assign to p1
	p1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}
	fmt.Println(p1)
	fmt.Println(p1.first)        // fields in person is accessible in secretAgent itself (promoted fields)
	fmt.Println(p1.person.first) // fields in person is accessible in secretAgent itself

	// avoid code pollution using anoynomous struct
	p2 := struct {
		first string
		last  string
	}{
		first: "Hanming",
		last:  "Zheng",
	}
	fmt.Println(p2)
}
