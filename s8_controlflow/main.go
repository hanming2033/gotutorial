package main

import (
	"fmt"
)

func main() {
	// // for loop, nested
	// for i := 0; i <= 5; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		fmt.Println(i, j)
	// 	}
	// }

	// // while equivalent
	// x := 0
	// for x < 10 {
	// 	fmt.Println(x)
	// 	x++
	// }

	// // break infinite loop
	// // infinite loop : good for server listener
	// y := 0
	// for {
	// 	fmt.Println(y)
	// 	if y > 10 {
	// 		break
	// 	}
	// 	y++
	// }

	// // next iteration using continue
	// z := 0
	// for z < 10 {
	// 	z++
	// 	if z < 5 {
	// 		// if z < 5, do not do anything
	// 		continue
	// 	}
	// 	fmt.Println(z)
	// }

	// // initialization statement
	// // initialize x and do a check, !!! scope of x is within the if
	// if x := 42; x > 2 {
	// 	fmt.Println("001")
	// }

	// switch case no fallthrough (default)
	x := 51531531
	switch x {
	default:
		fmt.Println("default value")
	case 1023:
		fmt.Println("x == 100")
	case -1, 1, 2:
		fmt.Println("-1, 1 or 2")
	case 120:
		fmt.Println("x == 120")
		fallthrough // once fallthrough, all below will be activated
	case 0:
		fmt.Println("either 0 or fallthrough from 120")
	}

	// if else using switch
	k := 0
	switch {
	case k == 0:
		fmt.Println("k == 0")
	case k > 0:
		fmt.Println("k > 0")
	case k < 0:
		fmt.Println("k > 0")
	}

}
