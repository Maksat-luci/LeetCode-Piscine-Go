package main

import (
	"os"

	"github.com/01-edu/z01"
)

func atoi(a string) int {
	var n byte
	var v int
	for i,_ := range a {
		d := a[i]
		if d >= 47 && d <= 57 {
			n = d - '0'
		}
		v *= 10
		v += int(n)
	}
	return v
}
func itoa(num int) string {
	str := ""
	result := ""
	if num == 0 {
		return "0"
	}else {
		for num != 0{
			str+= string(rune(num%10)+48)
			num /= 10
		}
	}
	temp := []rune(str)
	for i := len(temp)-1 ; i >=0 ; i--{
		result += string(temp[i])
	}
	return result
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
	var str string
	for i := 1; i != 10; i++ {
		str = itoa(i)
		str += " x "
		str += itoa(integer)
		str += " = "
		str += itoa(i * integer)
		str += "\n"
		for _, v := range str {
			z01.PrintRune(v)
	
		}
	}
}
