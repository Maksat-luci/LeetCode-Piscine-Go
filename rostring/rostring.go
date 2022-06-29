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
	word := []string{}
	x := ""
	b := ""
	for _, v := range args[0] {
		if v == ' ' && x != "" {
			word = append(word, string(x), " ")
			x = b
		}
		if v != ' ' {
			x += string(v)
		}
	}
	if x != "" {
		word = append(word, x, " ")
	}
	start, end := 0, 0
	sres := []string{}
	for i, v := range word {
		for _, n := range v {
			if n == 32 && end == 0 {
				end = i
				sres = word[start:end]
			}
		}
	}
	word = word[end+1:]
	for _, m := range sres {
		word = append(word, m)
	}
	for _, m := range word {
		for _, nn := range m {
			z01.PrintRune(nn)
		}
	}
	z01.PrintRune('\n')
}
