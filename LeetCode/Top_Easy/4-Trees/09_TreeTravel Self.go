import "fmt"


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func dfs(root *TreeNode) bool{

	var myStack []TreeNode
	tempNode = root
	myStack = append(myStack, root)

	for len(myStack) > 0 || tempNode != nil{

		tempNode = myStack[len(myStack) - 1] 
		myStack = myStack [:len(myStack) - 1]
		for tempNode != nil {
			if tempNode.Right != nil{
				myStack = append(myStack,tempNode.Right)
			}
			fmt.Println(tempNode.Val)
			tempNode = tempNode.Left
		}
	}
}

func inOrder(root *TreeNode, tree *[]int){

	if root == nil{
		return
	}

	inOrder(root.Left,tree)
	*tree = append(*tree, root.Val)
	inOrder(root.Right,tree)

}

func main(){

}


