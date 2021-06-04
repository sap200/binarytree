package tree

import (
	"fmt"

	"github.com/sap200/binarytree/node"
)

type BST struct {
	root *node.TreeNode
	size int
}

// Creates a new binary tree
func NewBST() *BST {
	return &BST{
		root: nil,
		size: 0,
	}
}

func CreateBST(el []node.Comparable) *BST {
	bst := NewBST()
	for i := 0; i < len(el); i++ {
		bst.Insert(el[i])
	}

	return bst
}

func (bst *BST) Search(e node.Comparable) bool {

	current := bst.root

	for current != nil {

		if e.CompareTo(current.Element) < 0 { // element lesser traverse left subtree
			current = current.LeftNode
		} else if e.CompareTo(current.Element) > 0 { // element greater traverse right subtree
			current = current.RightNode
		} else {
			// element found
			return true
		}
	}

	return false
}

// Inserts a node in binary tree
func (bst *BST) Insert(e node.Comparable) bool {

	if bst.GetRoot() == nil {
		bst.root = node.NewTreeNode(e)
		bst.size++
		return true
	}

	// finds the location of parent and inserts
	var parent *node.TreeNode
	current := bst.root

	for current != nil {
		if e.CompareTo(current.Element) < 0 {
			parent = current
			current = parent.LeftNode
		} else if e.CompareTo(current.Element) > 0 {
			parent = current
			current = parent.RightNode
		} else {
			// duplicate element found
			return false
		}
	}

	// parent found
	if e.CompareTo(parent.Element) < 0 { // e less than parent element
		parent.LeftNode = node.NewTreeNode(e)
	} else { // e greater than parent node
		parent.RightNode = node.NewTreeNode(e)
	}

	bst.size++
	return true
}

// delete an element
func (bst *BST) Delete(e node.Comparable) bool {
	// Locate the position of to be deleted node and its parent
	var parent *node.TreeNode
	current := bst.root

	for current != nil {
		if e.CompareTo(current.Element) < 0 {
			parent = current
			current = parent.LeftNode
		} else if e.CompareTo(current.Element) > 0 {
			parent = current
			current = parent.RightNode
		} else {
			// position of element found
			break
		}
	}

	// If current is nil then element doesn't exists
	if current == nil {
		return false
	}

	// case 1
	// Left subtree doesn't exists for current node
	if current.LeftNode == nil {
		// delete current node
		// attach right subtree to parent
		if parent == nil {
			bst.root = current.RightNode
		} else {
			if e.CompareTo(parent.Element) < 0 {
				// connect to left node of parent
				parent.LeftNode = current.RightNode
			} else {
				parent.RightNode = current.RightNode
			}
		}

	} else {
		// case 2
		// Find rightmost element
		parentOfRightMost := current
		rightMost := current.LeftNode

		for rightMost.RightNode != nil {
			parentOfRightMost = rightMost
			rightMost = parentOfRightMost.RightNode
		}

		// Replace current Element with rightmost
		current.Element = rightMost.Element

		// Delete Rightmost Element
		if parentOfRightMost.RightNode == rightMost {
			// then rightmost is not current
			parentOfRightMost.RightNode = rightMost.LeftNode
		} else {
			// current is rightmost
			parentOfRightMost.LeftNode = rightMost.LeftNode

		}
	}

	bst.size--
	return true

}

// Traversal methods

// inorder is a recursive helper function
func inorder(node *node.TreeNode) {
	if node == nil {
		return
	}

	inorder(node.LeftNode)
	fmt.Print(node.Element, " ")
	inorder(node.RightNode)

}

// It calls the inorder function for inorder traversal
func (b *BST) Inorder() {
	inorder(b.GetRoot())
}

// preorder is a helper function for preorder traversal
func preorder(node *node.TreeNode) {
	if node == nil {
		return
	}

	fmt.Print(node.Element, " ")
	preorder(node.LeftNode)
	preorder(node.RightNode)
}

// Preorder calls the preorder function
func (bst *BST) Preorder() {
	preorder(bst.GetRoot())
}

// postorder is a helper function for postorder traversal
func postorder(node *node.TreeNode) {
	if node == nil {
		return
	}

	preorder(node.LeftNode)
	preorder(node.RightNode)
	fmt.Print(node.Element, " ")
}

// Postorder calls the postorder function
func (bst *BST) Postorder() {
	postorder(bst.GetRoot())
}

// Utilities function

// Returns the size of binary tree
func (b *BST) GetSize() int {
	return b.size
}

// checks if the binary tree is empty returns true on not empty
func (b *BST) IsEmpty() bool {
	return b.GetSize() == 0
}

// clears the binary tree
func (b *BST) Clear() {
	b.size = 0
	b.root = nil
}

// returns the root of binary tree
func (bst *BST) GetRoot() *node.TreeNode {
	return bst.root
}

// Helper function for sorting
func sorted(node *node.TreeNode, list *[]node.Comparable) {
	if node == nil {
		return
	}

	sorted(node.LeftNode, list)
	*list = append(*list, node.Element)
	sorted(node.RightNode, list)

}

// Returns a list of sorted node
func (bst *BST) Sorted() []node.Comparable {
	list := &[]node.Comparable{}
	sorted(bst.GetRoot(), list)
	return *list
}

// Print Binary tree
func (tree *BST) Print() {
	// Traverse tree
	fmt.Print("Inorder (sorted): ")
	tree.Inorder()
	fmt.Print("\nPostorder: ")
	tree.Postorder()
	fmt.Print("\nPreorder: ")
	tree.Preorder()
	fmt.Print("\nThe number of nodes is ", tree.GetSize())
	fmt.Println()
}
