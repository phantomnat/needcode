package easy

import (
	"sort"
)

type KthLargest struct {
	Nums []int
	K    int
}

func NewKthLargest(k int, nums []int) KthLargest {
	out := KthLargest{
		Nums: nums,
		K:    k,
	}
	sort.Ints(out.Nums)
	return out
}

func (this *KthLargest) Add(val int) (out int) {
	defer func() {
		out = this.Nums[len(this.Nums)-this.K]
	}()

	if len(this.Nums) == 0 {
		this.Nums = []int{val}
		return
	} else if val < this.Nums[0] {
		this.Nums = append([]int{val}, this.Nums...)
		return
	}

	l := 0
	r := len(this.Nums)
	for l < r {
		m := (l + r) / 2
		if this.Nums[m] <= val {
			l = m + 1
		} else {
			r = m
		}
	}
	if l == len(this.Nums) {
		this.Nums = append(this.Nums, val)
	} else {
		this.Nums = append(this.Nums[:l+1], this.Nums[l:]...)
		this.Nums[l] = val
	}
	
	return out
}
