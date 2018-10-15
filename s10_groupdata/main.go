package main

import (
	"fmt"
)

func main() {
	// * Array (fixed size: [5]) => Do not use array, use slices
	// array in go in fixed width
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)

	// * Slices (variable size: [])
	// composite literal
	slice := []int{8, 9, 10}
	y := []int{1, 2, 3, 4} // initialize slice can also do with []int{}
	fmt.Println(y)
	fmt.Println("length:", len(y))  // length of slice
	fmt.Println("err...", y[4])     // index out of range
	fmt.Println(y[:2])              // slice a slice: index 0 to index 1
	fmt.Println(y[1:2])             // slice a slice: index 1 to index 1
	fmt.Println(y[3:])              // slice a slice: index 3 to end
	y1 := append(y, 5, 6, 7)        // add items to slice
	y2 := append(y1, slice...)      // spread slice into a slice
	y3 := append(y2[:2], y2[4:]...) // delete item/s between indexs using appended and spread
	fmt.Println(y3)

	madeSlice := make([]int, 10, 100) // create slice, initialize 10 values, this slice can store 100 int, once slice size > 100, capacity becomes 200
	fmt.Println(madeSlice[10])        // index out of range

	row1 := [][]string{
		[]string{"A1", "A2", "A3", "A4"},
		[]string{"A1", "A2", "A3", "A4"},
	} // multi dimensional array
	fmt.Println(row1)

	// * Map => key value pair, super fast and efficient
	m := map[string]int{
		"James":      32,
		"Moneypenny": 27,
	} // map[string]int is the type
	fmt.Println(m)
	fmt.Println(m["James"])   // get value from map
	fmt.Println(m["hanming"]) // zero value will be returned if key not found
	v, exist := m["hanming"]  // get value from map and check if key exist
	fmt.Println(v)
	fmt.Println(exist)
	if v, exist := m["hanming"]; exist { // print value only if key exists
		fmt.Println(v)
	}
	m["hanming"] = 31  // add item to map
	delete(m, "James") // remove item from map
	for k, v := range m {
		fmt.Println(k, v)
	}
}
