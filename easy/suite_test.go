package easy

import (
	"encoding/json"
	"math"
	"testing"
	"time"

	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/ginkgo/v2/dsl/core"
	"github.com/onsi/ginkgo/v2/dsl/decorators"
	g "github.com/onsi/gomega"
	"github.com/onsi/gomega/gmeasure"
	"github.com/stretchr/testify/assert"
)

var p = &Practice{}

func TestEasy(t *testing.T) {
	g.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "medium problems")
}

var _ = BeforeSuite(func() {
	p = &Practice{}
})

var _ = Describe("easy problems", func() {

	Describe("merge two sorted list", func() {
		type input struct {
			List1 *ListNode
			List2 *ListNode
		}
		var testCases = []struct {
			input  input
			expect *ListNode
		}{
			{
				input: input{
					List1: NewLinkedList([]int{1, 2, 3, 4, 5}),
					List2: NewLinkedList([]int{1,4,5}),
				},
				expect: NewLinkedList([]int{1,1,2,3,4,4,5,5}),
			},
			{
				input: input{
					List1: nil,
					List2: nil,
				},
				expect: nil,
			},
		}
		It("check", func() {
			for i := range testCases {
				tc := testCases[i]
				actual := p.MergeTwoSortedList(tc.input.List1, tc.input.List2)
				check2(i, tc.expect, actual, tc.input)
			}
		})
	})

	Describe("reverse linked list", func() {
		type input struct {
			Head *ListNode
		}

		var testCases = []struct {
			input  input
			expect *ListNode
		}{
			{
				input: input{
					Head: NewLinkedList([]int{1, 2, 3, 4, 5}),
				},
				expect: NewLinkedList([]int{5, 4, 3, 2, 1}),
			},
			{
				input: input{
					Head: nil,
				},
				expect: nil,
			},
		}
		It("check", func() {
			for i := range testCases {
				tc := testCases[i]
				actual := p.ReverseLinkedList(tc.input.Head)
				check2(i, tc.expect, actual, tc.input)
			}
		})
	})

	Describe("binary search", func() {
		type input struct {
			Nums   []int
			Target int
		}

		var testCases = []struct {
			input  input
			expect int
		}{
			{
				input: input{
					Nums:   []int{-1, 0, 3, 5, 9, 12},
					Target: 9,
				},
				expect: 4,
			},
			{
				input: input{
					Nums:   []int{-1, 0, 3, 5, 9, 12},
					Target: 2,
				},
				expect: -1,
			},
		}
		It("check", func() {
			for i := range testCases {
				tc := testCases[i]
				actual := p.BinarySearch(tc.input.Nums, tc.input.Target)
				check2(i, tc.expect, actual, tc.input)
			}
		})
	})

	Describe("valid parentheses", func() {
		type input struct {
			S string
		}

		var testCases = []struct {
			input  input
			expect bool
		}{
			{
				input: input{
					S: "(())",
				},
				expect: true,
			},
			{
				input: input{
					S: "([)]",
				},
				expect: false,
			},
		}
		It("check", func() {
			for i := range testCases {
				tc := testCases[i]
				actual := p.ValidParentheses(tc.input.S)
				check2(i, tc.expect, actual, tc.input)
			}
		})
	})

	Describe("climb stairs", func() {

		type input struct {
			N int
		}

		var testCases = []struct {
			input  input
			expect int
		}{
			{
				input: input{
					N: 1,
				},
				expect: 1,
			},
			{
				input: input{
					N: 2,
				},
				expect: 2,
			},
			{
				input: input{
					N: 3,
				},
				expect: 3,
			},
			{
				input: input{
					N: 4,
				},
				expect: 5,
			},
			{
				input: input{
					N: 5,
				},
				expect: 8,
			},
		}

		It("check", func() {
			for i := range testCases {
				tc := testCases[i]
				actual := p.ClimbStairs(tc.input.N)
				check2(i, tc.expect, actual, tc.input)
			}
		})

		XIt("benchmark", decorators.Serial, func() {
			expr := gmeasure.NewExperiment("climb stairs benchmark")
			ginkgo.AddReportEntry(expr.Name, expr)

			expr.SampleDuration("my climb stairs", func(_ int) {
				for i := range testCases {
					tc := testCases[i]
					p.ClimbStairs(tc.input.N)
				}
			}, gmeasure.SamplingConfig{Duration: 2 * time.Second})

			expr.SampleDuration("another climb stairs", func(_ int) {
				for i := range testCases {
					tc := testCases[i]
					p.ClimbStairs2(tc.input.N)
				}
			}, gmeasure.SamplingConfig{Duration: 2 * time.Second})
		})
	})
})

func check2(i int, expect, actual, input any) {
	raw, _ := json.Marshal(input)
	g.Expect(actual).WithOffset(1).To(g.Equal(expect), "test case %d with input %s", i, string(raw))
}

func TestContainsDuplicate(t *testing.T) {
	a := assert.New(t)

	testCases := []struct {
		input  []int
		expect bool
	}{
		{
			input:  []int{1, 2, 3, 1},
			expect: true,
		},
		{
			input:  []int{1, 2, 3, 4},
			expect: false,
		},
		{
			input:  []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expect: true,
		},
	}
	for i := range testCases {
		tc := testCases[i]
		actual := p.ContainsDuplicate(tc.input)
		raw, _ := json.MarshalIndent(tc.input, "", "  ")
		a.Equal(tc.expect, actual, "test case %d with input %s", i, string(raw))
	}
}

func TestTwoSum(t *testing.T) {
	a := assert.New(t)

	type input struct {
		Nums   []int
		Target int
	}
	testCases := []struct {
		input  input
		expect []int
	}{
		{
			input: input{
				Nums:   []int{2, 7, 11, 15},
				Target: 9,
			},
			expect: []int{0, 1},
		},
		{
			input: input{
				Nums:   []int{3, 2, 4},
				Target: 6,
			},
			expect: []int{1, 2},
		},
		{
			input: input{
				Nums:   []int{3, 3},
				Target: 6,
			},
			expect: []int{0, 1},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		raw, _ := json.MarshalIndent(tc.input, "", "  ")
		actual := p.TwoSum(tc.input.Nums, tc.input.Target)
		a.Equal(tc.expect, actual, "test case %d with input %s", i, string(raw))
	}
}

func TestIsAnagram(t *testing.T) {
	a := assert.New(t)
	type input struct {
		S string
		T string
	}
	testCases := []struct {
		input  input
		expect bool
	}{
		{
			input: input{
				S: "anagram",
				T: "nagaram",
			},
			expect: true,
		},
		{
			input: input{
				S: "rat",
				T: "car",
			},
			expect: false,
		},
	}
	for i := range testCases {
		tc := testCases[i]
		actual := p.IsAnagram(tc.input.S, tc.input.T)
		raw, _ := json.MarshalIndent(tc.input, "", "  ")
		a.Equal(tc.expect, actual, "test case %d with input %s", i, string(raw))
	}
}

func TestIsPalindrome(t *testing.T) {
	a := assert.New(t)
	type input struct {
		S string
	}
	testCases := []struct {
		input  input
		expect bool
	}{
		{
			input: input{
				S: "A man, a plan, a canal: Panama",
			},
			expect: true,
		},
		{
			input: input{
				S: "race a car",
			},
			expect: false,
		},
		{
			input: input{
				S: " ",
			},
			expect: true,
		},
	}
	for i := range testCases {
		tc := testCases[i]
		actual := p.IsPalindrome(tc.input.S)
		raw, _ := json.MarshalIndent(tc.input, "", "  ")
		a.Equal(tc.expect, actual, "test case %d with input %s", i, string(raw))
	}
}

func TestFindMinCostClimbingStairs(t *testing.T) {
	a := assert.New(t)
	type input struct {
		Cost []int
	}
	testCases := []struct {
		input  input
		expect int
	}{
		{
			input: input{
				Cost: []int{10, 15, 20},
			},
			expect: 15,
		},
		{
			input: input{
				Cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			},
			expect: 6,
		},
	}
	for i := range testCases {
		tc := testCases[i]
		actual := p.FindMinCostClimbingStairs(tc.input.Cost)
		check(a, i, tc.expect, actual, tc.input)
	}
}

func check(a *assert.Assertions, i int, expect, actual, input any) {
	raw, _ := json.Marshal(input)
	a.Equal(expect, actual, "test case %d with input %s", i, string(raw))
}

type isHappyNumberInput struct {
	N int
}

var isHappyNumberTestCases = []struct {
	input  isHappyNumberInput
	expect bool
}{
	{
		input: isHappyNumberInput{
			N: 19,
		},
		expect: true,
	},
	{
		input: isHappyNumberInput{
			N: 2,
		},
		expect: false,
	},
	{
		input: isHappyNumberInput{
			N: math.MaxInt32,
		},
		expect: false,
	},
}

func TestIsHappyNumber(t *testing.T) {
	a := assert.New(t)
	type input struct {
		N int
	}

	for i := range isHappyNumberTestCases {
		tc := isHappyNumberTestCases[i]
		actual := p.IsHappyNumber(tc.input.N)
		check(a, i, tc.expect, actual, tc.input)
	}
}

func BenchmarkIsHappyNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := range isHappyNumberTestCases {
			tc := isHappyNumberTestCases[i]
			p.IsHappyNumber(tc.input.N)
		}
	}
}

func TestPlusOne(t *testing.T) {
	a := assert.New(t)
	type input struct {
		Digits []int
	}
	testCases := []struct {
		input  input
		expect []int
	}{
		{
			input: input{
				Digits: []int{9},
			},
			expect: []int{1, 0},
		},
		{
			input: input{
				Digits: []int{1, 2, 3},
			},
			expect: []int{1, 2, 4},
		},
		{
			input: input{
				Digits: []int{4, 3, 2, 1},
			},
			expect: []int{4, 3, 2, 2},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		actual := p.PlusOne(tc.input.Digits)
		check(a, i, tc.expect, actual, tc.input)
	}
}

func TestBestTimeToBuyAndSellStock(t *testing.T) {
	a := assert.New(t)
	type input struct {
		Prices []int
	}
	testCases := []struct {
		input  input
		expect int
	}{
		{
			input: input{
				Prices: []int{7, 1, 5, 3, 6, 4},
			},
			expect: 5,
		},
		{
			input: input{
				Prices: []int{7, 6, 4, 3, 1},
			},
			expect: 0,
		},
	}
	for i := range testCases {
		tc := testCases[i]
		actual := p.MaximizeProfit(tc.input.Prices)
		check(a, i, tc.expect, actual, tc.input)
	}
}
