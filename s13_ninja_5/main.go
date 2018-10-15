package main

import (
	"fmt"
)

type person struct {
	first       string
	last        string
	favFlavours []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	fourWheel bool
	vehicle
}

type sedan struct {
	luxury bool
	vehicle
}

func main() {
	p1 := person{
		first:       "Hanming",
		last:        "Zheng",
		favFlavours: []string{"I", "do", "not", "like", "ice", "cream"},
	}
	fmt.Println(p1)
	p2 := person{
		first:       "Wenyang",
		last:        "Si",
		favFlavours: []string{"vanilla", "straw berry"},
	}
	fmt.Println(p2)
	// 2
	fmt.Println("2.....")
	m := map[string]person{
		p1.last: p1,
	}
	m[p2.last] = p2
	for k, v := range m {
		fmt.Println(k)
		fmt.Printf("\t%v\n", v.first)
		fmt.Printf("\t%v\n", v.last)
		for _, v2 := range v.favFlavours {
			fmt.Printf("\t\t%v\n", v2)
		}
	}
	// 3
	fmt.Println("3.....")
	truck1 := truck{
		fourWheel: true,
		vehicle: vehicle{
			color: "red",
			doors: 2,
		},
	}
	fmt.Println(truck1.color)
	fmt.Println(truck1.doors)
	fmt.Println(truck1.fourWheel)
	sedan1 := sedan{
		luxury: true,
		vehicle: vehicle{
			color: "blue",
			doors: 4,
		},
	}
	fmt.Println(sedan1.color)
	fmt.Println(sedan1.doors)
	fmt.Println(sedan1.luxury)

	wat := struct {
		name string
		age  int
	}{
		name: "hanming",
		age:  31,
	}

	fmt.Println(wat)
}
