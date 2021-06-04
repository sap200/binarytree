# Binary Tree

Binary tree is a package that helps creating binary tree of custom elements

Create a custom element 

**

type Element struct {
	x string
}

**

Let the custom element be comparable

**

func (e Element) CompareTo(el node.Comparable) int {
	if e.x < el.(Element).x {
		return -1
	} else if e.x == el.(Element).x {
		return 0
	} else {
		return 1
	}
}

**

Perform Operations

**

package main

import (
	"fmt"
	"github.com/sap200/binarytree/node"
	"github.com/sap200/binarytree/tree"
)

func main() {
	var bin *tree.BST
	bin = tree.NewBST()
	bin.Insert(Element{"George"})
}

**
