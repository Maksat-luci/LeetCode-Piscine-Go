package main

import (
	"fmt"
	"os"
)

func atoi(s string) int {
	var n int
	var v byte
	for i, _ := range s {
		d := s[i]
		if d >= 48 && d <= 57 {
			v = d - '0'
		}
		n *= 10
		n += int(v)
	}
	return n
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
	firstNum := atoi(args[0])
	secondNum := atoi(args[1])
	integer := []int{}
	integerNum2 := []int{}
	for i := 1; i <= firstNum; i++ {
		if firstNum%i == 0 {      
			integer = append(integer, i)
		}
	}
	for j := 1; j <= secondNum; j++ {
		if secondNum%j == 0 {
			integerNum2 = append(integerNum2, j)
		}
	}
	var final_result int
	for _, g := range integer {
		for _, f := range integerNum2 {
			if g == f {
				final_result = f
			}
		}
	}
	fmt.Println(final_result)
}
