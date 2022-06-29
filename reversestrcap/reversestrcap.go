package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	end := 0
	result := ""
	for i := 0; i < len(args); {
		for _, m := range args[i] {
			if m >= 65 && m <= 90 {
				result += string(m + 32)
				continue
			}
			result += string(m)
		}
		final_result := []byte(result)
		for j, _ := range final_result {
			if j+1 != len(final_result) && final_result[j+1] == 32 {
				end = j
				final_result[end] -= 32
			}
		}
		length := 0
		for range final_result {
			length++
		}
		length -= 1
		if final_result[length] != 32 {
			final_result[length] -= 32
		}
		last := ""
		for _, f := range final_result {
			last += string(f)
		}
		for _, l := range last {
			z01.PrintRune(rune(l))
		}
		z01.PrintRune('\n')
		result = ""
		i++
	}
}
