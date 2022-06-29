package main

import (
	"fmt"
	"strings"
)

func bubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
// Itoa
func Itoa(nbr int, b bool) string {
	str := ""
	if nbr == 0{
		return "0"
	}
 	for nbr != 0 {
		str += string(rune(nbr%10) + 48)
		nbr /= 10
	}
	result := ""
	if b {
		result = "-"
	}
	res := []rune(str)
	for i := len(res) - 1; i >= 0; i-- {
		result += string(res[i])
	}
	return result
}

//Atoi ...
func Atoi(s string) int {
	var n int
	var v byte
	for i, _ := range s {
		d := s[i]
		if d >= 48 && d <= 57 {
			v = d - '0'
		}
		n *= 10
		n += int(v)
	}
	if strings.Contains(s, "-") {
		n = n * -1
	}
	return n
}

func Split(s, sep string) []string {
	var x int
	result := []string{}
	if sep == "" && len(s) > 0 {
		for _, vs := range s {
			result = append(result, string(vs))
		}
		return result
	}
	for i := 0; i < len(s)-len(sep); i++ {
		if s[i:len(sep)-1] == sep {
			result = append(result, s[x:i])
			x = i + len(sep)
			if sep != "" {
				i = x - 1
			}
		}
	}
	result = append(result, s[x:len(s)-1])
	if sep == "" {
		result = result[1 : len(s)-1]
	}
	return result
}

// HighAndLow ...
// решить проблему с дроблением строки по символам
func HighAndLow(in string) string {
	if len(in) == 0 {
		return ""
	}
	temporaryVariable := []int{}
	Variable := strings.Split(in, " ")
	for _, m := range Variable {
		temporaryVariable = append(temporaryVariable, Atoi(m))
	}
	temporaryVariable = bubbleSort(temporaryVariable)
	res := ""
	fmt.Println(temporaryVariable)
	fmt.Println(Itoa(0,true))
	if temporaryVariable[len(temporaryVariable)-1] < 0 {
		cheker := true
		s := temporaryVariable[len(temporaryVariable)-1] * -1
		res += Itoa(s, cheker)
	} else {
		res += Itoa(temporaryVariable[len(temporaryVariable)-1], false)
	}
	res += " "
	if temporaryVariable[0] <= 0 {
		cheker := true
		v := temporaryVariable[0] * -1
		res += Itoa(v, cheker)
	} else {
		res += Itoa(temporaryVariable[0], false)
	}
	return res
}

func main() {
	fmt.Println(HighAndLow("0 0 0 0 0 0 -4 0"))

}
