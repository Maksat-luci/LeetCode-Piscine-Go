package main

import (
	"fmt"
)

func compare(a, b string) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i <= len(a)-1; i++ {
		for j := i + 1; j <= len(a); j++ {
			if f(a[i], a[j]) > 0 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	AdvancedSortWordArr(result, compare)

	fmt.Println(result)
}
