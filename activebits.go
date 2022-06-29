package student

func ActiveBits(n int) uint {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return uint(count)
}
