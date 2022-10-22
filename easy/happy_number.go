package arraynhashing

func (_ Practice) IsHappyNumber(n int) bool {
	mem := make(map[int]int)

	squareAndSum := func(v int) int {
		total := 0
		for v > 0 {
			total += (v % 10) * (v % 10)
			v = v / 10
		}
		return total
	}

	for {
		if n == 1 {
			return true
		}
		if _, ok := mem[n]; ok {
			return false
		}
		mem[n] = squareAndSum(n)
		n = mem[n]
	}
}
