package easy

func (p *Practice) ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	mem := make([]int, n+1)
	mem[n-1] = 1
	mem[n-2] = 2

	for i := n - 3; i >= 0; i-- {
		mem[i] = mem[i+1] + mem[i+2]
	}
	return mem[0]
}

func (p *Practice) ClimbStairs2(n int) int {
	if n < 4 {
		return n
	}
	first, second := 1, 2
	for i := 3; i <= n; i++ {
		second, first = first+second, second
	}
	return second
}
