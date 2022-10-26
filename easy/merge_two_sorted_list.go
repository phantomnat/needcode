package easy

func (*Practice) MergeTwoSortedList(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	var last *ListNode

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			if head == nil {
				head = list1
				last = list1
			} else {
				last.Next = list1
				last = last.Next
			}
			list1 = list1.Next
		} else {
			if head == nil {
				head = list2
				last = list2
			} else {
				last.Next = list2
				last = last.Next
			}
			list2 = list2.Next
		}
	}
	for list1 != nil {
		if head == nil {
			head = list1
			last = list1
		} else {
			last.Next = list1
			last = last.Next
		}
		list1 = list1.Next
	}
	for list2 != nil {
		if head == nil {
			head = list2
			last = list2
		} else {
			last.Next = list2
			last = last.Next
		}
		list2 = list2.Next
	}
	return head
}
