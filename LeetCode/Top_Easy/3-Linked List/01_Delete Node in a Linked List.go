package main

import(
	"fmt"
)

type ListNode struct{
	Val int
	Next *ListNode
}

func main(){
}

func deleteNode(node *ListNode) {
	node.Val, node.Next = node.Next.Val, node.Next.Next  
}