package kata

//kyu 6
func CountBits(number uint) int {
	var bitcount int
	for {
		bitcount += int(number % 2)

		if number <= 1 {
			break
		}
		number /= 2
	}
	return bitcount
}
