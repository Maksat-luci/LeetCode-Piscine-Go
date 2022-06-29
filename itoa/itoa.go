package main

import "fmt"

func itoa(nbr int) string {
	str := ""
	for nbr != 0 {
		str = string(rune(nbr%10)+48) + str
		nbr /=10
	}
	return str
}

func main() {
	s := 12345
	fmt.Println(itoa(s))
}
