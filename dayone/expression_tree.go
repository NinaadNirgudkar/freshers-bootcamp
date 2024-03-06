package dayone

import "fmt"

func Express() {
	expression := "a+b-c"
	expTree := ExpressionTree{[]*Node{nil}}
	for _, c := range expression {
		str := string(c)
		switch str {
		case "+", "-":
			//If a character is an operator pop two values
			//from the stack make them its child
			//and push the current node again.
			fmt.Println("operator")
			left := expTree.pop()
			right := expTree.pop()
			newNode := &Node{left: left, right: right, value: str}
			expTree.push(newNode)
		default:
			fmt.Println("operand")
			// If a character is an operand push that into the stack
			newNode := &Node{left: nil, right: nil, value: str}
			expTree.push(newNode)
		}
	}
}

type ExpressionTree struct {
	tree []*Node
}

func (e *ExpressionTree) pop() (node *Node) {
	treeSize := len(e.tree)
	if treeSize == 0 {
		return
	}
	//lastNode := e.tree[treeSize - 1]
	newTree := ExpressionTree{}
	lastNode := e.tree[treeSize-1]
	for _, n := range e.tree {
		if lastNode != n {
			newTree.push(n)
		}
	}
	e.tree = newTree.tree
	return
}
func (e *ExpressionTree) push(node *Node) {
	treeSize := len(e.tree)
	if treeSize != 0 {
		lastNode := e.tree[treeSize-1]
		lastNode.right = node
		node.left = lastNode
	}
	e.tree = append(e.tree, node)
}

type Node struct {
	value string
	left  *Node
	right *Node
}
