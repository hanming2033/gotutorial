package main

import (
	"fmt"
	"reflect"
)

type ownInt int

var x ownInt

var y int

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v\n", x)
	fmt.Println(reflect.TypeOf(x).Kind())
	y = int(x)
}
