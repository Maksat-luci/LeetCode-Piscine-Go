package main

import "fmt"

func main() {
	s := "Let there     be light"
	var b string
	var x string
	var f []string
	for _, d := range s {
		if d == ' ' && x != "" {
			f = append(f, x)
			x = b
		}
		if d != ' ' {
			x += string(d)
		}
	}
	if x != "" {
		f = append(f, x)
	}
	fmt.Println(f)
}
