package student

func Split(s, sep string) []string {
	var x int
	result := []string{}
	if sep == "" && len(s) > 0 {
		for _, vs := range s {
			result = append(result, string(vs))
		}
		return result
	}
	for i := 0; i < len(s)-len(sep)+1; i++ {
		if s[i:len(sep)+i] == sep {
			result = append(result, s[x:i])
			x = i + len(sep)
			if sep != "" {
				i = x - 1
			}
		}
	}
	result = append(result, s[x:len(s)])
	if sep == "" {
		result = result[1 : len(s)-1]
	}
	return result

}
