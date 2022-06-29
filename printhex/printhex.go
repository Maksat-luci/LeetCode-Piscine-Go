package main

import (
	"os"

	"github.com/01-edu/z01"
)

func bubbleSort(array []int) (l []int) {
	var spisok []int
	for i := len(array) - 1; i > -1; {
		spisok = append(spisok, array[i])
		i--
	}
	return spisok
}
func printERROR() {
	for _, n := range "ERROR" {
		z01.PrintRune(n)
	}
	z01.PrintRune('\n')
}
func printHex(s []string){
	for _,b := range s {
		for _,n := range b {
			z01.PrintRune(n)
		}
	}
	z01.PrintRune('\n')
}

func atoi(s string) (int,bool){
	var n byte
	var v int
	for i, _ := range s {
		d := s[i]
		if d >= 48 && d <= 57 {
			n = d - '0'
		} else {
			printERROR()
			return 0, true
		}
		v *= 10
		v += int(n)
	}
	return v,false
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
	number,boolss := atoi(args[0])
	if boolss == true{
		return
	}
	base16 := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	slise := []int{}
	system := len(base16)
	for number != 0 {
		slise = append(slise, number%system)
		number /= system
	}
	rev_array := bubbleSort(slise)
	final_result := []string{}
	for _,k := range rev_array {
		final_result = append(final_result, base16[k] )
	}
	printHex(final_result)
}
