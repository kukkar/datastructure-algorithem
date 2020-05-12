package binarytree

func FindMaxDepth(tree *TreeNode)int {

	if tree == nil {
		return 0
	}

	lheight := lheight(tree.LeftNode)
	rheight := rheight(tree.RightNode)

	leftMax := FindMaxDepth(tree.Left)
	rightMax := FindMaxDepth(tree.Right)
	return max(lheight+rheight,max(leftMax,rightMax))
}

func height(n *TreeNode)int {
    if n == nil {
        return 0
    }
    return 1 + max(height(n.LeftNode),height(n.RightNode) )
}

func max (a,b int)int {
    if a >=b {
        return a
    } 
    return b
}
