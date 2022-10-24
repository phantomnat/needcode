package easy

func (p *Practice) TwoSum(nums []int, target int) []int {
	mem := make(map[int]int)
	for i, no := range nums {
		if _, ok := mem[no]; ok {
			return []int{mem[no], i}
		}
		mem[target-no] = i
	}
	return nil
}
