package main

import (
	"fmt"
	"os"
)

func atoi(s string) int {
	var v int
	var n byte
	for i, _ := range s {
		d := s[i]
		if d >= 48 && d <= 57 {
			n = d - '0'
		}
		v *= 10
		v += int(n)
	}
	return v
}

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		return
	}
	num := atoi(args[0])
	sum := 0
	for i := 2; i <= num; i++ {
		if IsPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}

func IsPrime(num int) bool {
	counter := 0
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			counter++
		}
	}
	if counter == 2 {
		return true
	}
	return false
}
