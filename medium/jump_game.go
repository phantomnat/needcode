package medium

func (p *Practice) CanJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	} else if nums[0] == 0 {
		return false
	}

	last := len(nums) - 1

	for i := last - 1; i >= 0; i-- {
		if nums[i]+i < last {
			continue
		}
		last = i
	}

	return nums[0] >= last
}
