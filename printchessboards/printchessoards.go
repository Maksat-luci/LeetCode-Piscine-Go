package main

import (
	"os"

	"github.com/01-edu/z01"
)

func atoi(s string) int {
	var b int
	var x byte
	for i, _ := range s {
		d := s[i]
		if d >= 48 && d <= 57 {
			x = d - '0'
		}
		b *= 10
		b += int(x)
	}
	return b
}

func main() {
	args := os.Args[1:]
	length := 0
	for range args {
		length++
	}
	if length != 2 {
		return
	}
	x, y := atoi(args[0]), atoi(args[1])
	res := ""
	restart := x
	for i := y; i != 0; {
		if i%2 == 0 {
			for x != 0 {
				res += " "
				x--
				if x == 0 {
					break
				}
				res += "#"
				x--
			}
		} else {
			for x != 0 {
				res += "#"
				x--
				if x == 0 {
					break
				}
				res += " "
				x--
			}
		}
		if i > 0 {
			x = restart
		}
		res += "\n"
		i--
	}
	for _, v := range res {
		z01.PrintRune(v)
	}

}