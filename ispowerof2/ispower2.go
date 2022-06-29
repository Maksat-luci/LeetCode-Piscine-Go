package main

import (
	"os"
	"strconv"
)

func IsPowerOfTwo(nbr int) bool {
	if nbr <= 0 || nbr == 1 {
		return false
	}
	for nbr != 1 {
		if nbr%2 != 0 {
			return false
		}
		nbr /= 2
	}
	return true
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	nbr, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}
	result := IsPowerOfTwo(nbr)
	if result {
		os.Stdout.WriteString("true\n")
	} else {
		os.Stdout.WriteString("false\n")
	}
}
