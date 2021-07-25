type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

import "fmt"

type queue struct{
	data []*TreeNode
}

func (q *queue) pop() *TreeNode{
	if len(q.data) == 0{
		return nil
	}

	topNode := q.data[0]
	q.data = q.data[1:]

	return topNode
}

func (q *queue) peek() *TreeNode{
	return q.data[0]
}

func (q *queue) push(node *TreeNode){
	q.data = append(q.data, node)
}

func (q *queue) empty() bool{
	return len(q.data) == 0
}

func (q. *queue) qSize() int{
	return len(q.data)
}

func newQueue() *queue{
	q := queue{}
	q.data = make([]*TreeNode, 0, 0)
	return &q
}
