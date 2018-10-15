package main

import (
	"fmt"
)

func main() {
	// 2
	sliceInt := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Printf("%T\n", sliceInt)
	for _, v := range sliceInt {
		fmt.Println(v)
	}
	// 3
	fmt.Println("3.....")
	fmt.Println(sliceInt[:5])
	fmt.Println(sliceInt[5:])
	fmt.Println(sliceInt[2:7])
	fmt.Println(sliceInt[1:6])
	// 4
	fmt.Println("4.....")
	sliceInt4 := append(sliceInt, 52)
	fmt.Println(sliceInt4)
	sliceInt4_1 := append(sliceInt4, 53, 54, 55)
	fmt.Println(sliceInt4_1)
	dummy := []int{56, 57, 58, 59, 60}
	sliceInt4_2 := append(sliceInt4_1, dummy...)
	fmt.Println(sliceInt4_2)
	// 5
	fmt.Println("5.....")
	sliceInt5 := append(sliceInt[:3], sliceInt[6:]...)
	fmt.Println(sliceInt5)
	// 6
	fmt.Println("6.....")
	sliceCities := make([]string, 50, 50)
	sliceCities = []string{`Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`, `Connecticut`, `Delaware`, `Florida`, `Georgia`, `Hawaii`, `Idaho`, `Illinois`, `Indiana`, `Iowa`, `Kansas`, `Kentucky`, `Louisiana`, `Maine`, `Maryland`, `Massachusetts`, `Michigan`, `Minnesota`, `Mississippi`, `Missouri`, `Montana`, `Nebraska`, `Nevada`, `New Hampshire`, `New Jersey`, `New Mexico`, `New York`, `North Carolina`, `North Dakota`, `Ohio`, `Oklahoma`, `Oregon`, `Pennsylvania`, `Rhode Island`, `South Carolina`, `South Dakota`, `Tennessee`, `Texas`, `Utah`, `Vermont`, `Virginia`, `Washington`, `West Virginia`, `Wisconsin`, `Wyoming`}
	fmt.Println(len(sliceCities))
	fmt.Println(cap(sliceCities))
	for i := 0; i < len(sliceCities); i++ {
		fmt.Println(i, sliceCities[i])
	}
	// 7
	fmt.Println("7.....")
	stringsTable := [][]string{[]string{"James", "Bond", "Shaken, not stirred"}, []string{"Miss", "Moneypenny", "Helloooooo, James."}}
	for i, row := range stringsTable {
		fmt.Println("Row:", i)
		for x, col := range row {
			fmt.Printf("\t %v(%v)\n", col, x)
		}
	}
	// 8
	fmt.Println("8.....")
	m := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	// 9
	m["hanming"] = []string{"yaya", "doudou", "wenyan", "mengdi"}
	// 10
	delete(m, "no_dr")
	for k, v := range m {
		fmt.Println(k)
		for i, v := range v {
			fmt.Printf("\t %v: %v", i, v)
		}
		fmt.Println()
	}
}
