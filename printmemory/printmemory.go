package main

import "fmt"

func PrintMemory(b [10]byte) {
	for i, l := range b {
		fmt.Printf("%.2x", l)
		if (i+1)%4 != 0 && i != len(b)-1 {
			fmt.Print(" ")
		} else if (i+1)%4 == 0 {
			fmt.Print("\n")
		}
	}
	fmt.Println()
	for _, l := range b {
		if l >= 33 && l <= 126 {
			fmt.Print(string(l))
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
