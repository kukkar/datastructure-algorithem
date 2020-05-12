package binarytree

func FindMaxDepth(tree *TreeNode)int {

	if tree == nil {
		return 0
	}

	lheight := height(tree.LeftNode)
	rheight := height(tree.RightNode)

	leftMax := FindMaxDepth(tree.LeftNode)
	rightMax := FindMaxDepth(tree.RightNode)
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
