package main

import (
	"fmt"
	"os"
)

func EndSort(value []int) []int {
	result := []int{}
	for i := len(value) - 1; i >= 0; i-- {
		result = append(result, value[i])
	}
	return result
}

func ItoaBase(value int, base string) string {
	system := len(base)
	array := []int{}
	for value != 0 {
		array = append(array, value%system)
		value /= system
	}
	revArray := EndSort(array)
	last_array := []string{}
	for _, f := range base {
		last_array = append(last_array, string(f))
	}
	finalString := ""
	for _, nb := range revArray {
		if nb < 0 {
			nb = nb * -1
			finalString += last_array[nb]
		} else {
			finalString += last_array[nb]
		}
	}
	return finalString
}

func atoi(val string) int {
	var n int
	var v byte
	for i, _ := range val {
		d := val[i]
		if d >= 48 && d <= 57 {
			v = d - 48
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
	value := atoi(args[0])
	base := args[1]
	fmt.Println(ItoaBase(value, base))
}
