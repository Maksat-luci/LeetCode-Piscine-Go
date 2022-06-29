package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	length := 0
	result := []rune{}
	for range args {
		length++
	}
	if length != 1 {
		return
	}
	word := args[0]
	for _, n := range word {
		if n >= 97 && n <= 122 {
			if n >= 'o' {
				n -= 13
				result = append(result, n)
			} else {
				result = append(result, n+13)
			}
		} else if n <= 90 && n >= 65 {
			if n >= 'O' {
				n -= 13
				result = append(result, n)
			} else {
				result = append(result, n+13)
			}
		} else {
			result = append(result, n)
		}
	}
	for _, b := range result {
		z01.PrintRune(b)
	}
	z01.PrintRune('\n')

}
