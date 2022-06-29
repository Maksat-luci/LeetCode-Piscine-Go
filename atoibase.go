package student

import (
	"github.com/01-edu/z01"
)

func Chek(base string) (kek bool) {
	c := 0
	cc := []string{}
	chekingbool := false
	for range base {
		c++
	}
	if c < 2 {
		return chekingbool
	}
	for _, l := range base {
		cc = append(cc, string(l))
	}
	for _, k := range base {
		if k == 43 || k == 45 {
			return chekingbool
		}
	}
	for i := 0; i < c-1; i++ {
		for j := i + 1; j < c; j++ {
			if cc[i] == cc[j] {
				return chekingbool
			}
		}
	}
	chekingbool = true
	return chekingbool
}

func PrintZero() {
	z01.PrintRune('0')
}

func RecursivePowers(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power == 1 {
		return nb
	}
	return nb * RecursivePowers(nb, power-1)
}

func AtoiBase(s string, base string) int {
	chekingresult := Chek(base)
	var integer []int
	system := len(base)
	var final_result int
	if chekingresult == false {
		return 0
	}
	for _, m := range s {
		for i, x := range base {
			if m == x {
				integer = append(integer, i)
			}
		}
	}
	mnozhitel := []int{}
	for i := len(integer) - 1; i > -1; i-- {
		mnozhitel = append(mnozhitel, RecursivePowers(system, i))
	}
	for i := 0; i < len(integer); i++ {
		final_result += integer[i] * mnozhitel[i]
	}
	return final_result
}
