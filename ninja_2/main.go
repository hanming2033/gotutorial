package main

import (
	"fmt"
)

func main() {
	a := 1234
	fmt.Printf("%d\n", a)
	fmt.Printf("%b\n", a)
	fmt.Printf("%#x\n", a)

	const typed int = 10
	const untyped = 10
	fmt.Println(typed, untyped)

	p := 12
	fmt.Printf("%d\t%b\t%#x\n", p, p, p)
	o := p << 1
	fmt.Printf("%d\t%b\t%#x\n", o, o, o)

	const (
		year1 = 2019 + iota
		year2
		year3
		year4
	)
	fmt.Printf("%d %d %d %d", year1, year2, year3, year4)
}
