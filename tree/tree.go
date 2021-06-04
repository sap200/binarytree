package tree

import "github.com/sap200/binarytree/node"

type Tree interface {

	// Search searches for an element in tree
	// returns true on success
	Search(element node.Comparable) bool

	// Insert inserts element into tree
	// returns true on successful insertion
	Insert(element node.Comparable) bool

	// Delete deletes an element from tree
	// returns true on success
	Delete(element node.Comparable) bool

	// Basic Traversal methods for tree
	Preorder()
	Inorder()
	Postorder()

	// Basic Utility functions

	// Returns number of TreeNodes on tree
	GetSize() int

	// Clears the tree
	Clear()

	// Returns true if the tree is empty
	IsEmpty() bool
}
