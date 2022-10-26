package medium

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/ginkgo/v2/dsl/core"
	g "github.com/onsi/gomega"

	"neetcode/test"
)

var (
	p *Practice
)

func TestMedium(t *testing.T) {
	g.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "medium problems")
}

var _ = BeforeSuite(func() {
	p = &Practice{}
})

var _ = Describe("medium problems", func() {
	It("jump game", func() {
		type input struct {
			Nums []int
		}
		testCases := []struct {
			input  input
			expect bool
		}{
			{
				input: input{
					Nums: []int{2, 3, 1, 1, 4},
				},
				expect: true,
			},
			{
				input: input{
					Nums: []int{3, 2, 1, 0, 4},
				},
				expect: false,
			},
			{
				input: input{
					Nums: []int{3, 2, 2, 0, 4},
				},
				expect: true,
			},
		}
		for i := range testCases {
			tc := testCases[i]
			actual := p.CanJump(tc.input.Nums)
			test.Check(i, tc.expect, actual, tc.input)
		}
	})

	XIt("insert interval", func() {
		type input struct {
			Interval    [][]int
			NewInterval []int
		}
		testCases := []struct {
			input  input
			expect [][]int
		}{
			{
				input: input{
					Interval:    [][]int{{1, 3}, {6, 9}},
					NewInterval: []int{2, 5},
				},
				expect: [][]int{{1, 5}, {6, 9}},
			},
			{
				input: input{
					Interval:    [][]int{{1, 3}, {6, 9}},
					NewInterval: []int{4, 5},
				},
				expect: [][]int{{1, 3}, {4, 5}, {6, 9}},
			},
			{
				input: input{
					Interval:    [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
					NewInterval: []int{4, 8},
				},
				expect: [][]int{{1, 2}, {3, 10}, {12, 16}},
			},
		}
		for i := range testCases {
			tc := testCases[i]
			actual := p.InsertInterval(tc.input.Interval, tc.input.NewInterval)
			test.Check(i, tc.expect, actual, tc.input)
		}
	})
})
