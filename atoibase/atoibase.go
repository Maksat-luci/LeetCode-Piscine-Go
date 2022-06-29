package main

import "fmt"

func recursivePowers(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power == 1 {
		return nb
	}
	return nb * recursivePowers(nb, power-1)
}

func AtoiBase(s string, base string) int {
	if chek(base) {
		return 0
	}
	indexes := []int{}
	system := len(base)
	for _, n := range s {
		for i, b := range base {
			if n == b {
				indexes = append(indexes, i)
			}
		}
	}
	mnozhitel := []int{}
	for i := len(indexes)-1; i >-1; i-- {
		mnozhitel = append(mnozhitel, recursivePowers(system, i))
	}
	res := 0
	for i := 0; i < len(indexes); i++ {
		res += indexes[i] * mnozhitel[i]
	}
	return res
}

func chek(base string) bool {
	for _, n := range base {
		if n == 43 || n == 45 {
			return true
		}
	}
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base)-1; j++ {
			if base[i] == base[j] {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}
