package main

type ListNode struct {
    Val int
    Next *ListNode
}
 

func main(){

}

// Two Pointer
func hasCycle(head *ListNode) bool {
    if head == nil{
        return false
    }
    nHead := head.Next
    
    for head != nil && nHead != nil && nHead.Next != nil {
        if head == nHead{
            return true
        }
        head = head.Next
        nHead = nHead.Next
        nHead = nHead.Next
    }
    return false
}