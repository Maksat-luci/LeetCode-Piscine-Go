package main

import (
	"fmt"
	"os"
)

func matchBrackets(exp string) bool {
	runes := []rune(exp)
	var opened []rune
	for _, l := range runes {
		if l == '{' || l == '[' || l == '(' {
			opened = append(opened, l)
		}
		if (l == '}') && opened[len(opened)-1] == '{' {
			opened = opened[:len(opened)-1]
			continue
		}
		if (l == ']') && opened[len(opened)-1] == '[' {
			opened = opened[:len(opened)-1]
			continue
		}
		if (l == ')') && opened[len(opened)-1] == '(' {
			opened = opened[:len(opened)-1]
			continue
		}
		if (l == '}') && opened[len(opened)-1] != '{' {
			return false
		}
		if (l == ']') && opened[len(opened)-1] != '[' {
			return false
		}
		if (l == ')') && opened[len(opened)-1] != '(' {
			return false
		}

	}
	return true
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println()
	} else {
		for _, v := range os.Args[1:] {
			if matchBrackets(v) {
				fmt.Println("OK")
			} else {
				fmt.Println("Error")
			}
		}
	}
}