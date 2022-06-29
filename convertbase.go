package student

import (
	"github.com/01-edu/z01"
)

func NbrToStrRec2(n, dot int64) string {
	if 10 > n && n > -1*10 {
		return string('0' + n*dot)
	}
	return NbrToStrRec2(n/10, dot) + string('0'+(n%10)*dot)
}

func NbrToStr2(n int64) string {
	dot := int64(1)
	res := ""
	if n == 0 {
		return "0"
	}
	if n < 0 {
		dot *= -1
		res += "-"
	}
	return res + NbrToStrRec2(n, dot)
}

func Chek2(base string) (kek bool) {
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

func RecursivePower2(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if power == 1 {
		return nb
	}
	return nb * RecursivePower(nb, power-1)
}

func AtoiBase2(s string, base string) int {
	chekingresult := Chek2(base)
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
	var mnozhitel []int
	for i := len(integer) - 1; i > -1; i-- {
		mnozhitel = append(mnozhitel, RecursivePower(system, i))
	}
	for i := 0; i < len(integer); i++ {
		final_result += integer[i] * mnozhitel[i]
	}
	return final_result
}

func printNV() {
	NV := "NV"
	for _, k := range NV {
		z01.PrintRune(k)
	}
}

func chekingforNBRBase(base string) (kek bool) {
	c := 0
	cc := []string{}
	chekingbool := false
	for range base {
		c++
	}
	if c < 2 {
		PrintNV()
		return chekingbool
	}
	for _, l := range base {
		cc = append(cc, string(l))
	}
	for _, k := range base {
		if k == 43 || k == 45 {
			PrintNV()
			return chekingbool
		}
	}
	for i := 0; i < c-1; i++ {
		for j := i + 1; j < c; j++ {
			if cc[i] == cc[j] {
				PrintNV()
				return chekingbool
			}
		}
	}
	chekingbool = true
	return chekingbool
}

func bubbleSort(array []int) (l []int) {
	var spisok []int

	for i := len(array) - 1; i > -1; {
		spisok = append(spisok, array[i])
		i--
	}
	return spisok
}

func printConsole(result []string, booling bool) string {
	var lastString string
	if booling == true {
		lastString += "-"
		for _, j := range result {
			for _, gg := range j {
				lastString += string(gg)
			}
		}
	}
	if booling == false {
		for _, j := range result {
			for _, gg := range j {
				lastString += string(gg)
			}
		}
	}
	return lastString
}

func printNbrBase(nbr int, base string) string {
	booling := false
	if nbr < 0 {
		nbr = nbr * -1
		booling = true
	}
	chekingresult := chekingforNBRBase(base)
	result := []int{}
	if chekingresult == false {
		return "\n"
	}
	system := len(base)
	for nbr != 0 {
		result = append(result, nbr%system)
		nbr /= system
	}
	rev_array := bubbleSort(result)
	last_array := []string{}
	for _, k := range base {
		last_array = append(last_array, string(k))
	}
	final_array := []string{}
	for _, nv := range rev_array {
		final_array = append(final_array, last_array[nv])
	}
	ress := printConsole(final_array, booling)
	return ress
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	printNBRbass := AtoiBase2(nbr, baseFrom)
	result := printNbrBase(printNBRbass, baseTo)

	return result
}
