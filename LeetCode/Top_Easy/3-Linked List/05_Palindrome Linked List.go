package main


type ListNode struct {
    Val int
    Next *ListNode
}
 

func main(){

}


func isPalindrome(head *ListNode) bool {
    
    lSize := 0
    iterNode := head

    for ; iterNode != nil; lSize++ {
        iterNode = iterNode.Next
    }
        
    iterNode = head
    
    for i := 0; i < lSize/2; i++ {
        iterNode = iterNode.Next
    }

    secList := reverseList (iterNode)
    iterNode = head

    for iterNode != nil && secList != nil{
        if secList.Val != iterNode.Val{
            return false
        }        
        iterNode = iterNode.Next
        secList = secList.Next
    }

    return true


}



// Recur
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