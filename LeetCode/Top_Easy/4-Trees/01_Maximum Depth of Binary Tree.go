package main

import "fmt"

type NodeH struct{
    root *TreeNode
    height int
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func main(){

}

// Iter
func maxDepth(root *TreeNode) int {
    
    if root == nil{
        return 0
    }

    queue := make([]NodeH, 0)
    
    var temp NodeH
    temp.root = root
    temp.height = 1

    queue = append(queue, temp)
    // cHeight := 1
    maxs := 1

    for len(queue)>0{
        nNode := queue[0]
        fmt.Println(nNode.root.Val)
        queue := queue[1:]

        var ntemp NodeH
        ntemp.height = nNode.height+1

        if nNode.root.Left != nil{
            ntemp.root = nNode.root.Left
            queue = append(queue,ntemp)
            if ntemp.height>maxs{
                maxs = ntemp.height
            }
        }

        if nNode.root.Right != nil{
            ntemp.root = nNode.root.Right
            queue = append(queue,ntemp)
            if ntemp.height>maxs{
                maxs = ntemp.height
            }
        }
    }

    return maxs
}


// Recur
// func maxDepth(root *TreeNode) int {

//     if root == nil{
//         return 0
//     }

//     left := maxDepth(root.Left) + 1
//     right := maxDepth(root.Right) + 1

//     if left > right {
//         return left
//     }
//     return right
// }
