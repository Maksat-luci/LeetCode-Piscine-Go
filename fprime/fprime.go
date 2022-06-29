package main

import (
	"fmt"
	"os"
)

func atoi(s string) int {
	var v int
	var b byte
	for i, _ := range s {
		d := s[i]
		if d >= 48 && d <= 57 {
			b = d - '0'
		} else {
			return 0
		}
		v *= 10
		v += int(b)
	}
	return v
}

func Fprime(value int) {
	if value == 1 {
		return
	}
	divisionIterator := 2
	for value > 1 {
		if value%divisionIterator == 0 {
			fmt.Print(divisionIterator)
			value = value / divisionIterator
			if value > 1 {
				fmt.Print("*")
			}
			divisionIterator--
		}
		divisionIterator++
	}
	fmt.Println()
}

func main() {
	args := os.Args[1:]
	length := 0
	for range args {
		length++
	}
	if length != 1 {
		return
	}
	integer := atoi(args[0])
	if integer == 0 {
		return
	}
	Fprime(integer)
}
