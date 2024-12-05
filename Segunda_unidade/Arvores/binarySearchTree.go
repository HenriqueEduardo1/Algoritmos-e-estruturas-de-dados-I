package main

import "fmt"

type BSTNode struct {
	left  *BSTNode
	val   int
	right *BSTNode
}

func createNode(val int) *BSTNode {
	return &BSTNode{val: val}
}

func (node *BSTNode) Add(val int) {
	if val < node.val {
		if node.left != nil {
			node.left.Add(val)
		} else {
			node.left = createNode(val)
		}
	} else {
		if node.right != nil {
			node.right.Add(val)
		} else {
			node.right = createNode(val)
		}
	}
}

func (node *BSTNode) Min() int {
	if node.left == nil {
		return node.val
	}
	return node.left.Min()
}

func (node *BSTNode) Max() int {
	if node.right == nil {
		return node.val
	}
	return node.right.Max()
}

func (node *BSTNode) Height() int {
	hl := 0
	hr := 0
	hMax := 0

	if node.left == nil && node.right == nil {
		return 0
	}

	if node.left != nil {
		hl = node.left.Height()
	}

	if node.right != nil {
		hr = node.right.Height()
	}

	if hl > hr {
		hMax = hl
	} else {
		hMax = hr
	}

	return hMax + 1
}

func (node *BSTNode) Search(val int) bool {
	if node == nil {
		return false
	}
	if val < node.val {
		return node.left.Search(val)
	} else if val > node.val {
		return node.right.Search(val)
	}
	return true
}

func (node *BSTNode) InOrderNav() {
	if node.left != nil {
		node.left.InOrderNav()
	}
	fmt.Println(node.val)
	if node.right != nil {
		node.right.InOrderNav()
	}
}

func main() {
	// bst := createNode(10)
	// bst.Add(5)
	// bst.Add(15)
	// bst.Add(8)
	// bst.Add(3)
	// bst.Add(12)
	// bst.Add(20)
	// bst.Add(13)
	// bst.Add(7)
	// bst.Add(9)
	// bst.Add(11)

	// bst.InOrderNav()

	// fmt.Println(bst.Search(13))
	// fmt.Println(bst.Search(14))
	// fmt.Println(bst.Search(7))
	// fmt.Println(bst.Search(20))
	// fmt.Println(bst.Min())
	// fmt.Println(bst.Max())

	bst2 := createNode(10)
	bst2.Add(5)
	bst2.Add(20)
	bst2.Add(15)
	bst2.Add(16)

	bst2.InOrderNav()

	fmt.Println("Altura:", bst2.Height())

}
