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
	bin.Insert(Element{"Michael"})
	bin.Insert(Element{"Tom"})
	bin.Insert(Element{"Adam"})
	bin.Insert(Element{"Jones"})
	bin.Insert(Element{"Peter"})
	bin.Insert(Element{"Daniel"})

	bin.Print()
	bin.Delete(Element{"George"})
	fmt.Println()
	bin.Print()
	bin.Delete(Element{"Michael"})
	fmt.Println()
	bin.Print()
	bin.Delete(Element{"Adam"})
	fmt.Println()
	bin.Print()
	bin.Delete(Element{"Jones"})
	fmt.Println()
	bin.Print()
	bin.Delete(Element{"Tom"})
	fmt.Println()
	bin.Print()
	bin.Delete(Element{"Peter"})
	fmt.Println()
	bin.Print()
	bin.Delete(Element{"Daniel"})
	fmt.Println()
	bin.Print()

}

type Element struct {
	x string
}

func (e Element) CompareTo(el node.Comparable) int {
	if e.x < el.(Element).x {
		return -1
	} else if e.x == el.(Element).x {
		return 0
	} else {
		return 1
	}
}
