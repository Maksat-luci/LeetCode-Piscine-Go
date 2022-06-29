package main

import "fmt"

func IsAnagram(str1, str2 string) bool {
	count := 0
	for _, n := range str1 {
		for _, g := range str2 {
			if n == g {
				count++
				break
			}
		}
	}
	if count == len(str1) {
		return true
	}
	return false
}

func main() {
	test1 := IsAnagram("listen", "silent")
	fmt.Println(test1)

	test2 := IsAnagram("alem", "school")
	fmt.Println(test2)

	test3 := IsAnagram("neat", "a net")
	fmt.Println(test3)

	test4 := IsAnagram("anna madrigal", "a man and a girl")
	fmt.Println(test4)
}
