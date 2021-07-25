func levelOrder(root *TreeNode) [][]int{
	result := make([][]int, 0, 0)

	if root == nil{
		return [][]int{}
	}

	BFSQ := newQueue()
	BFSQ.push(root)

	for !BFSQ.empty(){
		var currSize int = BFSQ.qSize()
		currNodes := make([]int, 0, 0)

		for i := 0; i < currSize; i++{
			var topNode *TreeNode = BFSQ.pop()
			currNodes = append(currNodes, topNode.Val)
			if topNode.Left != nil {
				BFSQ.push(topNode.Left)
			}
			if topNode.Right != nil{
				BFSQ.push(topNode.Right)
			}
		}

		result = append(result, currNodes)
	}
}
