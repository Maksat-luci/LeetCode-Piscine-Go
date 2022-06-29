package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	length := 0
	for range args {
		length++
	}
	if length != 1 {
		return
	}
	word := args[0]
	res := []rune{}
	for _, m := range word {
		if m >= 97 && m <= 122 {
			res = append(res, 122-(m-97))
		} else if m >= 65 && m <= 90 {
			res = append(res, 90-(m-65))
		} else {
			res = append(res, m)
		}

	}
	for _, m := range res {
		z01.PrintRune(m)
	}
	z01.PrintRune('\n')
}
