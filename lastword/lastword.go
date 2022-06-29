package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	length := 0
	result := ""
	fl := false
	for range args {
		length++
	}
	if length != 1 {
		return
	}
	word := args[0]
	for i := len(word) - 1; i > 0; i-- {
		if word[i-1] != ' ' && fl == false {
			result = string(word[:i])
			fl = true
		}
	}
	for i, v := range result {
		if v == 32 {
			result = result[i+1:]
		}
	}
	for _, c := range result {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
