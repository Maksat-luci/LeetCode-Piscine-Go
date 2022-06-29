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
	integer := 0
	word := args[0]
	alphabit := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	alphabit2 := []rune{'A', 'B', 'C', 'E', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	for _, s := range word {
		if s <= 122 && s >= 97 {
			for i, c := range alphabit {
				if s == c {
					integer = i + 1
					l := 0
					for l != integer {
						l++
						z01.PrintRune(c)
					}
				}
			}
		}
		if s >= 65 && s <= 90 {
			for i, c := range alphabit2 {
				if s == c {
					integer = i + 1
					l := 0
					for l != integer {
						l++
						z01.PrintRune(c)

					}
				}
			}

		}
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')

}
