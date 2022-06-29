package main

import (
	"fmt"
	"os"
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
	word := args[0]
	count := 0
	for _, j := range word {
		if j == 'U' {
			count += 1
		} else if j == 'D' {
			count -= 1
		} else if j == 'R' {
			count += 2
		} else if j == 'L' {
			count -= 2
		}
	}
	if count == 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

}
