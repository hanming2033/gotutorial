package main

import (
	"fmt"
	"strings"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// any type with method speak is human
// a value can be of more than 1 type
type human interface {
	speak()
}

func main() {
	// // * everything in Go is PASS BY VALUE
	// foo()
	// bar("Hanming")
	// fmt.Println(woo("Hanming"))
	// // initialize two identifiers with results from function
	// formalName, isOwner := doubleReturn("Hanming", "Zheng")
	// fmt.Println(formalName)
	// fmt.Println(isOwner)
	// // invoking variadic function
	// fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8))

	// xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	// sumThese := sum(xi...)
	// fmt.Println(sumThese)
	// sumZeros := sum()
	// fmt.Println(sumZeros)

	// // defer function execution
	// defer cleanUpFn() // fooo() is executed right before the wrapping function is about to end
	// doSomething()

	// sa1 := secretAgent{
	// 	person: person{
	// 		first: "Hanming",
	// 		last:  "Zheng",
	// 	},
	// 	ltk: true,
	// }
	// sa1.speak() // My name is Hanming
	// // sa1 is of type secretAgent, because it has speak method, it is also of type human
	// sayHuman(sa1) // I am a human main.secretAgent

	// p1 := person{
	// 	first: "John",
	// 	last:  "Doe",
	// }
	// p1.speak()
	// sayHuman(p1) // I am a human main.person

	// // anonymous function
	// func() {
	// 	fmt.Println("Anonymous func ran")
	// }()
	// func(x int) {
	// 	fmt.Println("Anonymous func ran", x)
	// }(42)

	// // function expression
	// f := func(x int) {
	// 	fmt.Println("function expression", x)
	// }
	// f(42)

	// give1 is a function produced by giveFunc
	give1 := giveFunc(1)
	// give1 will return 1 on execution
	fmt.Println(give1())

	// invoking call back function
	printMulti(2, multi2)
}

// function that run a function(2nd arg) passing in int(1st arg)
func printMulti(v int, function func(x int) int) {
	fmt.Println(function(v))
}

// function that returns int
func multi2(x int) int {
	return x * 2
}

// return a function from a function
func giveFunc(x int) func() int {
	// giveFunc returns function which returns int [func() int]
	return func() int {
		return x
	}
}

func sayHuman(h human) {
	// sayHuman is general
	// use switch and assertion/casting to drill down to specifics
	switch h.(type) { // .(type)
	case secretAgent:
		fmt.Println("I am a human", h.(secretAgent).first)
	case person:
		fmt.Println("I am a human", h.(person).first)
	}
}

// method are attached to struct using receiver (sa secretAgent)
func (sa secretAgent) kill(name string) {
	fmt.Println("Agent", sa.last, "killed", name)
}

func (sa secretAgent) speak() {
	fmt.Println("My name is", sa.first)
}
func (p person) speak() {
	fmt.Println("My name is", p.first)
}

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("Hello", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo -", s)
}

// function that return 2 results
func doubleReturn(fn, ln string) (string, bool) {
	formalName := fmt.Sprint(fn, " ", ln)
	var isAdmin bool
	if strings.Contains(strings.ToLower(fn), "hanming") {
		isAdmin = true
	}
	return formalName, isAdmin
}

// variadic function: takes any number of arguments
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// defer function
func cleanUpFn() {
	fmt.Println("fooo")
}

func doSomething() {
	fmt.Println("barr")
}
