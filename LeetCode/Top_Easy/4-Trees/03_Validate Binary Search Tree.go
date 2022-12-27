import "fmt"




type TreeNode struct {
	Val int
	Left TreeNode
	Right TreeNode
}


func main(){




}

// Range Method 
func isValidBST(root *TreeNode) bool {

	return validate(root, nil, nil)

}

func validate(root *TreeNode, LeftLimit *TreeNode, RightLimit *TreeNode){

	if root == nil{
		return true
	}

	if (LeftLimit !=  nil && LeftLimit.Val >= root.Val) || (RightLimit != nil && RightLimit.Val <= root.Val){
		return false
	}

	return validate(root.Right, root, RightLimit) && validate(root.LeftLimit, LeftLimit, root)
}



// inorder then Check if Sort
func isValidBST(root *TreeNode) bool {

	var tree []int
	inOrder(root,&tree)
	treeSize := len(tree)
	for i := 1; i < treeSize; i++ {
		if tree[i] <= tree[i-1]{
			return false
		}
	}
	return true
}
	
func inOrder(root *TreeNode, tree *[]int){

	if root == nil{
		return
	}

	inOrder(root.Left,tree)
	*tree = append(*tree, root.Val)
	inOrder(root.Right,tree)

}