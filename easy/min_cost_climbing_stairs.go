package arraynhashing

func (p *Practice) FindMinCostClimbingStairs(cost []int) int {
	n := len(cost)
	mem := make([]int, n)
	mem[n-1] = cost[n-1]
	mem[n-2] = cost[n-2]

	min := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}

	for i := n - 3; i >= 0; i-- {
		oneStep := cost[i] + mem[i+1]
		twoSteps := cost[i] + mem[i+2]
		mem[i] = min(oneStep, twoSteps)
	}

	return min(mem[0], mem[1])
}
