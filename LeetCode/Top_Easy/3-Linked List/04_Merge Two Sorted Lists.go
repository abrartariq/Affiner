package main


type ListNode struct {
    Val int
    Next *ListNode
}
 

func main(){

}

// Iter
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    
    if list1 == nil{
        return list2
    }else if list2 == nil{
        return list1
    }

    if list1.Val > list2.Val{
        list1, list2 = list2, list1
    }

    head := list1
    var nextNode *ListNode = new(ListNode);

    for list1 != nil && list2 != nil{
        if list1.Val <= list2.Val{
            nextNode.Next = list1
            list1 = list1.Next 
        }else {
            nextNode.Next = list2
            list2 = list2.Next 
        }
        nextNode = nextNode.Next 
    }

    if list1 == nil{
        nextNode.Next = list2
    }else{
        nextNode.Next = list1
    }

    return head
}

// Recur
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    
    if list1 == nil{
        return list2
    }else if list2 == nil{
        return list1
    }

    if list1.Val <= list2.Val{
        list1.Next := mergeTwoLists(list1.Next,list2)
        return list1
    }else{
        list2.Next := mergeTwoLists(list1,list2.Next)
        return list2
    }

}
