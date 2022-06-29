package main

import (
	"os"

	"github.com/01-edu/z01"
)

func ReverseSlize(res []rune) (s []rune) {
	ss := res
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
	return ss
}

func PrintConsole(result []string) {
	for _, text := range result {
		for _, word := range text {
			z01.PrintRune(word)
		}
	}
	z01.PrintRune('\n')
}

func main() {
	args := os.Args[1:]
	resq := []rune{}
	svvg := []string{}
	c := 0
	l := 0
	nbr := 1
	for range args {
		c++
	}
	if c != 1 {
		z01.PrintRune('\n')
	}
	for i, v := range args {
		if i == nbr {
			svvg = append(svvg, " ")
			nbr++
		}
		for _, k := range v {
			svvg = append(svvg, string(k))
		}
	}
	vowels := []string{"a", "A", "e", "E", "i", "I", "o", "O", "U", "u"}
	for _, x := range args {
		for _, v := range x {
			for _, xx := range vowels {
				if string(v) == xx {
					resq = append(resq, v)
				}
			}
		}
	}
	res := ReverseSlize(resq)
	for i, x := range svvg {
		for _, xx := range vowels {
			if xx == string(x) {
				svvg[i] = string(res[l])
				l++
			}
		}
	}

	PrintConsole(svvg)
}
