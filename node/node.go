// package node
package node

// An interface that ensures that all elements are comparable
type Comparable interface {
	CompareTo(e Comparable) int
}

// The basic node of a tree
type TreeNode struct {
	Element   Comparable
	LeftNode  *TreeNode
	RightNode *TreeNode
}

// Create a new node with given element
func NewTreeNode(e Comparable) *TreeNode {
	return &TreeNode{
		Element:   e,
		LeftNode:  nil,
		RightNode: nil,
	}
}
