
# Binary Tree

Binary tree is a package that helps creating binary tree of custom elements

## Install

Initiate a project

Type 

```
go get github.com/sap200/binarytree

```

## Usage

Create a custom element 

```go

type Element struct {
	x string
}

```

Let the custom element be comparable

```go

func (e Element) CompareTo(el node.Comparable) int {
	if e.x < el.(Element).x {
		return -1
	} else if e.x == el.(Element).x {
		return 0
	} else {
		return 1
	}
}

```

Perform Operations

```go

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

```
For more details check [example.go](example.go)
