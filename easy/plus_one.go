package easy

func (p *Practice) PlusOne(digits []int) []int {
	add := 1

	for i := len(digits) - 1; add > 0 && i >= 0; i-- {
		v := digits[i] + 1
		add, digits[i] = v/10, v%10
	}

	if add == 1 {
		result := make([]int, len(digits)+1)
		result[0] = 1
		copy(result[1:], digits)
		return result
	}

	return digits
}
