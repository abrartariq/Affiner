import "fmt"


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// Recur
func isSymmetric(root *TreeNode) bool {
	if root == nil{
		return true
	}
	return validate(root.Left, root.Right)
}

func validate(LeftNode *TreeNode, RightNode *TreeNode) bool {
	if LeftNode == nil && RightNode == nil{
		return true
	}else if (LeftNode != nil && RightNode == nil) || (LeftNode == nil && RightNode != nil) || (LeftNode.Val != RightNode.Val){
		return false
	}else{
		return validate(LeftNode.Left,RightNode.Right) && validate(LeftNode.Right,RightNode.Left)
	}
}


// Iter BFS
func isSymmetric(root *TreeNode) bool {
	if root == nil || root.Left == nil && root.Right == nil{
		return true
	}
	return parallelBFS(root.Left, root.Right)
}

// Single Queue BFS
func parallelBFS(LeftNode *TreeNode, RightNode *TreeNode) bool{

	queue := []*TreeNode{LeftNode,RightNode} 
	
	for len(queue) > 0{

		tmpLeft := queue[0]
		tmpRight := queue[1]
		queue = queue[2:]

		if tmpLeft == nil && tmpRight == nil{
			continue	
		}else if tmpLeft == nil || tmpRight == nil || tmpLeft.Val != tmpRight.Val{
            return false
        }
		queue = append(queue,tmpLeft.Left,tmpRight.Right,tmpLeft.Right,tmpRight.Left)
	}
	return true
}


// Iter DFS
func isSymmetric(root *TreeNode) bool {
	if root == nil || root.Left == nil && root.Right == nil{
		return true
	}else if (root.Left == nil && root.Right != nil) ||  (root.Left != nil && root.Right == nil) || root.Left.Val != root.Right.Val{
		return false
	}
	return parallelDFS(root.Left, root.Right)
}

func parallelDFS(LeftNode *TreeNode, RightNode *TreeNode) bool{

	var myStackL []*TreeNode
	var myStackR []*TreeNode
	tempNodeL := LeftNode
	tempNodeR := RightNode
	myStackL = append(myStackL, LeftNode)
	myStackR = append(myStackR, RightNode)

	for len(myStackL) > 0 || tempNodeL != nil {

		tempNodeL = myStackL[len(myStackL) - 1] 
		myStackL = myStackL[:len(myStackL) - 1]

		tempNodeR = myStackR[len(myStackR) - 1] 
		myStackR = myStackR[:len(myStackR) - 1]

		for tempNodeL != nil {
			if tempNodeL.Val != tempNodeR.Val{
				return false
			}else if tempNodeL.Right != nil && tempNodeR.Left != nil{
				myStackL = append(myStackL,tempNodeL.Right)
				myStackR = append(myStackR,tempNodeR.Left)
			}
			if (tempNodeL.Right != nil && tempNodeR.Left == nil) || (tempNodeL.Right == nil && tempNodeR.Left != nil) || (tempNodeL.Left != nil && tempNodeR.Right == nil) || (tempNodeL.Left == nil && tempNodeR.Right != nil){
				return false
			}

			tempNodeL = tempNodeL.Left
			tempNodeR = tempNodeR.Right

		}

		if tempNodeR != nil {
			return false
		}
	}
	return true
}

