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
	if length != 2 {
		return
	}
	str := ""
	for _, z := range args[0] {
		for _, n := range args[1] {
			if z == n {
				str += string(z)
				break
			}
		}
	}
	laststr := ""
	for _, k := range str {
		cheker := true
		for i := 0; i < len(laststr)-1; i++ {
			if laststr[i] == byte(k) {
				cheker = false
				break
			}
		}
		if cheker {
			laststr += string(k)
		}
	}
	for _, l := range laststr {
		z01.PrintRune(l)
	}
	z01.PrintRune('\n')
}
