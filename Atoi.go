package student

func Atoi(s string) int {
	var n int
	var v byte
	cheker := false
	if s == ""{
		return 0
	}
	x := []byte(s)
	for i, c := range x {
		if len(s) <= 1 {
			if c < 46 || c > 58 {
				return 0
			}
		}
		if i != 0 && c < 46 || c > 58 {
			return 0
		}
	}
	for i, _ := range s {
		if i >= 20 {
			return 0 
		}
		d := s[i]
		if '0' <= d && d <= '9' {
			v = d - '0'
		} else if 'a' <= d && d <= 'z' {
			return 0
		} else if 'A' <= d && d <= 'Z' {
			return 0
		} else {
			n = 0
		}
		n *= 10
		n += int(v)
	}
	if x[0] >= 43 && x[0] <= 45 && x[1] >= 43 && x[1] <= 45 {
		return 0
	}
	if x[0] == '-' && x[1] != '-' {
		n = n * -1
	}
	for _, word := range x {
		if word == 32 {
			return 0
		}
	}
	if x[0] == '-' {
		cheker = true
	} 
	if cheker == true && int(n) > 0 && len(x) >= 19{
		return 0 
	}
	if cheker == false && int(n) < 0 && len(x) >=19{
		return 0
	}
	return int(n)
}
