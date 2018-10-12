package main

import (
	"fmt"
)

// create own type
type ownInt int

var b ownInt
var a int

// a != ownInt

func main() {
	a = 42
	b = 10
	// main.ownInt cannot assign to int
	/*a = b*/
	// int(main.ownInt) => int : type conversion
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
