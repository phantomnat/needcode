package types

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNodeFromInts(arr ...int) *TreeNode {
	var intp = make([]*int, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		intp = append(intp, &arr[i])
	}
	return NewTreeNode(intp)
}

func NewTreeNode(arr []*int) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}
	head := &TreeNode{Val: *arr[0]}

	var currents []*TreeNode = []*TreeNode{head}
	var nexts []*TreeNode
	var idx = 1

	for len(currents) > 0 && idx < len(arr) {
		node := currents[0]
		currents = currents[1:]

		func() {
			if node == nil {
				idx += 2
				return
			}

			if idx >= len(arr) {
				return
			}
			if arr[idx] != nil {
				node.Left = &TreeNode{Val: *arr[idx]}
			}
			nexts = append(nexts, node.Left)
			idx++

			if idx >= len(arr) {
				return
			}
			if arr[idx] != nil {
				node.Right = &TreeNode{Val: *arr[idx]}
			}
			nexts = append(nexts, node.Right)
			idx++
		}()

		if len(currents) == 0 && len(nexts) > 0 {
			currents = nexts
			nexts = nil
		}
	}

	return head
}
