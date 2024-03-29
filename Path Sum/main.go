package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }

    if root.Left == nil && root.Right == nil {
        return root.Val == targetSum
    }

    leftPath := hasPathSum(root.Left, targetSum - root.Val)
    rightPath := hasPathSum(root.Right, targetSum - root.Val)

    return leftPath || rightPath
}