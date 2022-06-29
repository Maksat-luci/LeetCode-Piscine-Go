package main

import (
	"fmt"
)

func Nauuo(plus, negative, vopros int) string {
	if plus > negative && plus > negative+vopros {
		return "+"
	} else if negative > plus && negative > plus+vopros {
		return "-"
	} else if plus == negative && plus +vopros == negative && negative + vopros == plus {
		return "0"
	}
	return "?"
}

func main() {
	fmt.Println(Nauuo(50, 43, 20))
	fmt.Println(Nauuo(13, 13, 0))
	fmt.Println(Nauuo(10, 9, 0))
	fmt.Println(Nauuo(5, 9, 2))
	fmt.Println(Nauuo(50, 100, 50))
}
