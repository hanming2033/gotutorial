package main

import (
	"fmt"
)

var x = 90.0

func main() {
	hello := "Hello"
	fmt.Println(hello)
	// w `` literal string, keeps everything inside
	s := `Hanming says "Hello"`
	// convert string to slice of bytes
	fmt.Println([]byte(s))

	// for loop print the utf8 char one by one
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v ", string(s[i]))
	}
	fmt.Println()
	// for range loop print the utf8 char one by one
	for _, v := range s {
		fmt.Printf("%v ", string(v))
	}
	fmt.Println()
	fmt.Println([]byte("世界"))

	// constants
	x = 80.09
}
