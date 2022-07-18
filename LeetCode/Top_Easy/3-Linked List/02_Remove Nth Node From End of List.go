package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head == nil {
		return nil
	} else if head.Next == nil && n > 0 {
		return nil
	} else {

		nthVal := head
		for i := 1; i < n && nthVal.Next != nil; i++ {
			nthVal = nthVal.Next
		}

		iterVal := head
		prev := head

		for nthVal.Next != nil {
			nthVal = nthVal.Next
			prev = iterVal
			iterVal = iterVal.Next
		}

		if iterVal == prev {
			return head.Next
		} else {
			prev.Next = iterVal.Next
			return head
		}
	}
}
