package dayone

import "fmt"

func Express() {
	expression := "a+b-c"
	for _, c := range expression {
		str := string(c)
		switch str {
		case "+", "-":
			fmt.Println("operator")
		default:
			fmt.Println("operand")

		}
	}
}

type ExpressionTree struct {
	tree []*Node
}

func (e *ExpressionTree) pop() {

}
func (e *ExpressionTree) push() {

}

type Node struct {
	value []byte
	left  *Node
	right *Node
}
