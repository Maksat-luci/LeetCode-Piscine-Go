package main

import "fmt"

func isPrime(x int) bool {
	count := 0
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			count++
		}
	}
	if count == 2 {
		return true
	}
	return false
}

func PriorPrime(x int) int {
	res := 0
	for i := 2; i <= x; i++ {
		if isPrime(i) {
			res += i
		}
	}
	return res

}

func main() {
	fmt.Println(PriorPrime(14))
}
