package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func max(a int, b int) int {
    if a > b {
        return a
    }

    return b
}

func helper(root *TreeNode, answer *int) int {
    if root == nil {
        return 0
    }

    maxPathLeft := max(helper(root.Left, answer), 0)
    maxPathRight := max(helper(root.Right, answer), 0)
    *answer = max(*answer, maxPathLeft + maxPathRight + root.Val)
    return max(maxPathLeft, maxPathRight) + root.Val
}

func maxPathSum(root *TreeNode) int {
    var answer int = root.Val
    helper(root, &answer)
    return answer
}
