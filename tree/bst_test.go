package tree

import (
	"fmt"
	"testing"

	"github.com/sap200/binarytree/node"
)

type Element struct {
	x int
}

type Element1 struct {
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

func (e Element1) CompareTo(el node.Comparable) int {
	if e.x < el.(Element1).x {
		return -1
	} else if e.x == el.(Element1).x {
		return 0
	} else {
		return 1
	}
}

func TestNewBST(t *testing.T) {
	bst := NewBST()
	if bst == nil {
		t.Error("Expected address got nil")
	}

	if bst.GetRoot() != nil {
		t.Error("Expected nil got", bst.GetRoot(), "\n")
		fmt.Println()
	}

	if bst.GetSize() != 0 {
		t.Fatal("Expected 0, got", bst.GetSize(), "\n")
	}
}

func TestCreateBST(t *testing.T) {
	e := []node.Comparable{
		Element{2},
		Element{4},
		Element{3},
		Element{1},
		Element{8},
		Element{5},
		Element{7},
	}

	answer := []int{1, 2, 3, 4, 5, 7, 8}

	bst := CreateBST(e)

	x := bst.Sorted()

	for i, el := range x {
		if el.(Element).x != answer[i] {
			t.Errorf("Expected %d, Got %d", i+1, el.(Element).x)
		}
	}
}

func TestSearch(t *testing.T) {

	e := []node.Comparable{
		Element{2},
		Element{4},
		Element{3},
		Element{1},
		Element{8},
		Element{5},
		Element{7},
	}

	bst := CreateBST(e)

	cases := []struct {
		element int
		result  bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, true},
		{5, true},
		{6, false},
		{7, true},
		{9, false},
		{-1, false},
	}

	for _, el := range cases {
		if bst.Search(Element{el.element}) != el.result {
			t.Errorf("Expected %v, Got %v", el.result, bst.Search(Element{el.element}))
		}
	}
}

func TestDelete(t *testing.T) {
	e := []node.Comparable{
		Element1{"George"},
		Element1{"Michael"},
		Element1{"Tom"},
		Element1{"Adam"},
		Element1{"Jones"},
		Element1{"Peter"},
		Element1{"Daniel"},
	}

	bst := CreateBST(e)

	cases := []struct {
		element string
	}{
		{"George"},
		{"Michael"},
		{"Tom"},
		{"Adam"},
		{"Jones"},
		{"Peter"},
		{"Daniel"},
	}

	for _, el := range cases {
		bst.Delete(Element1{el.element})
		list := bst.Sorted()
		for _, elmnt := range list {
			if elmnt.(Element1).x == el.element {
				t.Errorf("Unable to delete %s", el.element)
			}
		}
	}

	if !bst.IsEmpty() {
		fmt.Println(bst)
		t.Errorf("Expected true, Got %v ---- size %d", bst.IsEmpty(), bst.GetSize())
	}
}

func TestInsert(t *testing.T) {
	bst := NewBST()

	cases := []struct {
		element string
	}{
		{"George"},
		{"Michael"},
		{"Tom"},
		{"Adam"},
		{"Jones"},
		{"Peter"},
		{"Daniel"},
	}

	for _, el := range cases {
		bst.Insert(Element1{el.element})
		// check if element is found
		list := bst.Sorted()
		notPresent := true
		for _, elmnt := range list {
			if elmnt.(Element1).x == el.element {
				notPresent = false
			}
		}

		if notPresent {
			t.Errorf("Unable to Insert %v", el.element)
		}
	}

}

// Examples

func ExampleCreateBST() {
	e := []node.Comparable{
		Element{2},
		Element{4},
		Element{3},
		Element{1},
		Element{8},
		Element{5},
		Element{7},
	}

	bst := CreateBST(e)
	fmt.Println(bst)
	fmt.Println(bst)
}

func ExampleBST_Insert() {
	var bin *BST
	bin = NewBST()
	bin.Insert(Element1{"George"})
	bin.Insert(Element1{"Michael"})
	bin.Insert(Element1{"Tom"})
	bin.Insert(Element1{"Adam"})
	bin.Insert(Element1{"Jones"})
	bin.Insert(Element1{"Peter"})
	bin.Insert(Element1{"Daniel"})
	fmt.Println(bin)
}

func ExampleBST_Search() {
	var bin *BST
	bin = NewBST()
	bin.Insert(Element1{"George"})
	bin.Insert(Element1{"Michael"})
	bin.Insert(Element1{"Tom"})
	bin.Insert(Element1{"Adam"})
	bin.Insert(Element1{"Jones"})
	bin.Insert(Element1{"Peter"})
	bin.Insert(Element1{"Daniel"})

	fmt.Println("Exists?", bin.Search(Element1{"George"}))
	fmt.Println("Exists?", bin.Search(Element1{"Fawad"}))
}

func ExampleBST_Delete() {
	var bin *BST
	bin = NewBST()
	bin.Insert(Element1{"George"})
	bin.Insert(Element1{"Michael"})
	bin.Insert(Element1{"Tom"})
	bin.Insert(Element1{"Adam"})
	bin.Insert(Element1{"Jones"})
	bin.Insert(Element1{"Peter"})
	bin.Insert(Element1{"Daniel"})

	fmt.Println("before:")
	bin.Print()
	bin.Delete(Element1{"George"})
	fmt.Println("after:")
	bin.Print()
}
