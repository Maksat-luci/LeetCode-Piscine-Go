package student

func SplitWhiteSpaces(s string) []string {
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
	return f
}
