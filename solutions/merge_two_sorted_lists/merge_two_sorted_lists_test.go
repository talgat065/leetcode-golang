package solutions

import (
	. "leetcode-golang/structures"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	t.Run("Merge two sorted lists", func(t *testing.T) {
		l1 := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode {
					Val: 4,
				},
			},
		}
		l2 := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode {
					Val: 4,
				},
			},
		}

		list := MergeTwoLists(l1, l2)

		s := []int{1, 1, 2, 3, 4}
		for i := 0; i < len(s); i++ {
			if list.Val != s[i] {
				t.FailNow()
			}
			list = list.Next
		}
	})
}
