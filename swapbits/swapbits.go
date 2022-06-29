package main

import "fmt"

func Swapbits(octet byte) byte {
	return octet<<4 + octet>>4
}

func main() {
	fmt.Println(231)
}
