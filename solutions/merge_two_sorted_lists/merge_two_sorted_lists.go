package merge_two_sorted_lists

import (
	. "leetcode-golang/structures"
)

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	head := ListNode{}

	if list1.Val < list2.Val {
		head.Val = list1.Val
		head.Next = MergeTwoLists(list1.Next, list2)
	} else {
		head.Val = list2.Val
		head.Next = MergeTwoLists(list1, list2.Next)
	}

	return &head
}
