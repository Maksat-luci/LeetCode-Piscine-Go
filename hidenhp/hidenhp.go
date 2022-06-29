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
	if length != 2 {
		return
	}
	s1 := args[0]
	s2 := args[1]
	resq := ""
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				resq += string(s1[i])
				if i < len(s1)-1 {
					i++
				} else {
					break
				}

			}
		}
	}
	if resq == s1 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
