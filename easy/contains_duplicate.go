package easy

func (p *Practice) ContainsDuplicate(nums []int) bool {
	mem := make(map[int]struct{})
	for _, no := range nums {
		if _, ok := mem[no]; ok {
			return true
		}
		mem[no] = struct{}{}
	}
	return false
}
