package main

import (
	"github.com/01-edu/z01"
	// "os"
)

func printConsole(str string) {
	for _, n := range str {
		z01.PrintRune(n)
	}
	z01.PrintRune('\n')
}

func main() {
	args := []string{"zpadinton", "paqefwtdjetyiytjneytjoeyjnejeyj"}
	length := 0
	for range args {
		length++
	}
	if length != 2 {
		return
	}
	str := ""
	for _, n := range args {
		str += n
	}
	result := ""
	for _, n := range str {
		cheker := true
		for i := 0; i < len(result); i++ {
			if result[i] == byte(n) {
				cheker = false
				break
			}
		}
		if cheker {
			result += string(n)
		}
	}
	printConsole(result)
}
