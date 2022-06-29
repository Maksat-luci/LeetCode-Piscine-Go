package main

import (
	"os"
)

func PrintConsole(str string) {
	os.Stdout.WriteString(str)
	os.Stdout.WriteString("\n")
	os.Stdout.Close()
}

func NbrToStrRec(n, dot int64) string {
	if 10 > n && n > -1*10 {
		return string('0' + n*dot)
	}
	return NbrToStrRec(n/10, dot) + string('0'+(n%10)*dot)
}

func NbrToStr(n int64) string {
	dot := int64(1)
	res := ""
	if n == 0 {
		return "0"
	}
	if n < 0 {
		dot *= -1
		res += "-"
	}
	return res + NbrToStrRec(n, dot)
}

func Atoi(nbr string) int64 {
	var res int64 = 0
	var sign int64 = 1
	hex := true
	value := ""
	chek := true
	hacker := true
	length := 0
	for range nbr {
		length++
	}
	if nbr[0] == '-' {
		nbr = nbr[1:]
		sign *= -1
		chek = false
		hacker = false
	} else if nbr[0] == '+' {
		nbr = nbr[1:]
	}
	for i, m := range nbr {
		if m != 48 {
			nbr = nbr[i:]
			hex = false
			break
		}
	}
	lan := 0
	for range nbr {
		lan++
	}
	for _, digit := range nbr {
		tmp := int64(digit-'0') * sign
		res = res*10 + tmp
	}
	if chek == false && res > 0 {
		PrintConsole("0")
	}
	if lan > 19 && hex == false {
		PrintConsole("0")
	}
	if hacker == false {
		value += "-"
		value += nbr
	} else {
		value += nbr
	}
	cheker := NbrToStr(res)
	// cheker = cheker[1:]
	for i := 0; i != len(cheker); i++ {
		if value[i] == cheker[i] {
			continue
		} else {
			PrintConsole("0")
			break
		}
	}
	return res
}

func main() {
	args := os.Args[1:]
	argsCount := 0

	for range args {
		argsCount++
	}
	if argsCount != 3 {
		return
	}
	x, y := args[0], args[2]
	for _, n := range x {
		if n != 45 && n != 43 {
			if n <= 47 || n >= 58 {
				PrintConsole("0")
				return
			}
		}
	}
	for _, s := range y {
		if s != 45 && s != 43 {
			if s <= 47 || s >= 58 {
				PrintConsole("0")
				return
			}
		}
	}
	if args[1] == "-" && 0 > Atoi(args[2]) {
		PrintConsole(NbrToStr(Atoi(args[0]) - Atoi(args[2])))
		return
	}
	funcsArr := []func(string, string){Plus, Deduct, Devide, Multiply, Mod}
	operators := []string{"+", "-", "/", "*", "%"}
	for i, val := range operators {
		if val == args[1] {
			funcsArr[i](args[0], args[2])
			return
		}
	}

	PrintConsole("0")
}
func itoa(nbr int64) string {
	str := ""
	for nbr != 0 {
		str += string(rune(nbr%10) + 48)
		nbr /= 10
	}
	result := ""
	res := []rune(str)
	for i := len(res) - 1; i >= 0; i-- {
		result += string(res[i])
	}
	return result
}

func Plus(a, b string) {
	aa := Atoi(a)
	integer := aa
	bb := Atoi(b)
	lnd := len(itoa(aa))
	missResult := aa + bb
	if (missResult < integer || missResult > integer) && lnd >= 19 {
		PrintConsole("0")
	}

	PrintConsole(NbrToStr(aa + bb))
}

func Deduct(a, b string) {
	lnd := len(a)
	aa := Atoi(a)
	integer := aa
	bb := Atoi(b)
	missResult := aa - bb
	if missResult > integer && lnd >= 19 {
		PrintConsole("0")
	}
	PrintConsole(NbrToStr(aa - bb))
}

func Devide(a, b string) {
	bb := Atoi(b)
	if bb == 0 {
		PrintConsole("No division by 0")
		return
	}
	aa := Atoi(a)
	if bb == -9223372036854775808 && aa == -1 {
		PrintConsole("0")
	}
	if aa == -9223372036854775808 && bb == -1 {
		PrintConsole("0")
	}

	PrintConsole(NbrToStr(aa / bb))
}
func Multiply(a, b string) {
	aa := Atoi(a)
	bb := Atoi(b)
	if bb == 0 || aa == 0 {
		PrintConsole("0")
	}
	if aa == -9223372036854775808 && bb == -1 {
		PrintConsole("0")
	}
	if bb == -9223372036854775808 && aa == -1 {
		PrintConsole("0")
	}
	if (aa > 4611686018427387903 || aa < -4611686018427387904) && (bb > 1 || bb < -1) {
		PrintConsole("0")
	}
	if (bb > 4611686018427387903 || bb < -4611686018427387904) && (aa > 1 || aa < -1) {
		PrintConsole("0")
	}
	f := aa * bb
	x := f / aa
	if x != bb {
		PrintConsole("0")
	}
	PrintConsole(NbrToStr(aa * bb))

}

func Mod(a, b string) {
	bb := Atoi(b)
	if bb == 0 {
		PrintConsole("No modulo by 0")
		return
	}
	aa := Atoi(a)
	PrintConsole(NbrToStr(aa % bb))
}
