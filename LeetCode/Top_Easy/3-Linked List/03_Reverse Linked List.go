package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {


}

// Iter
func reverseList(head *ListNode) *ListNode {
    
    // 1 or 0 Case
    if head == nil || head.Next == nil {
    	
   }else{
    	var prevNode *ListNode := head
    	var currNode *ListNode := head.Next
    	var nextNode *ListNode := head.Next.Next

    	prevNode.Next = nil

    	for nextNode != nil{

    		currNode.Next = prevNode
    		prevNode = currNode
    		currNode = nextNode
    		nextNode = nextNode.Next
    	
    	}

    	currNode.Next = prevNode
    	head = currNode
    }

    return head
}

// recursive
func reverseList(head *ListNode) *ListNode {
    
    // 1 or 0 Case
    if head == nil || head.Next == nil{
    	return head
	}
	newNode := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

    return newNode
}