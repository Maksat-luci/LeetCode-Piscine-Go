package student

import (
	"github.com/01-edu/z01"
)

func PrintNV() {
	NV := "NV"
	for _, k := range NV {
		z01.PrintRune(k)
	}
}

func ChekingforNBRBase(base string) (kek bool) {
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

func BubbleSort(array []int) (l []int) {
	var spisok []int
	for i := len(array) - 1; i > -1; {
		spisok = append(spisok, array[i])
		i--
	}
	return spisok
}

func PrintConsole(result []string, booling bool) {
	if booling == true {
		z01.PrintRune('-')
		for _, j := range result {
			for _, gg := range j {
				z01.PrintRune(gg)
			}
		}
	}
	if booling == false {
		for _, j := range result {
			for _, gg := range j {
				z01.PrintRune(gg)
			}
		}
	}
}

func PrintNbrBase(nbr int, base string) {
	booling := false
	if nbr < 0 {
		nbr = nbr * -1
		booling = true
	}
	chekingresult := ChekingforNBRBase(base)
	result := []int{}
	if chekingresult == false {
		return
	}
	system := len(base)
	for nbr != 0 {
		result = append(result, nbr%system)
		nbr /= system
	}
	rev_array := BubbleSort(result)
	last_array := []string{}
	for _, k := range base {
		last_array = append(last_array, string(k))
	}
	final_array := []string{}
	for _, nv := range rev_array {
		if nv < 0 {
			nv = nv * -1
			final_array = append(final_array, last_array[nv])
		} else {
			final_array = append(final_array, last_array[nv])
		}
	}
	PrintConsole(final_array, booling)
}
