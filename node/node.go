// package node
package node

type Comparable interface {
	CompareTo(e Comparable) int
}

type TreeNode struct {
	Element   Comparable
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func NewTreeNode(e Comparable) *TreeNode {
	return &TreeNode{
		Element:   e,
		LeftNode:  nil,
		RightNode: nil,
	}
}
