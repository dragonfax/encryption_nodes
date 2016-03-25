package main

import "fmt"

type Operation int

const (
	PLUS Operation = iota
	MINUS
	MULT
	DIV
	SHIFT_LEFT
	SHIFT_RIGHT

	INPUT
	OUTPUT
)

type Node struct {
	// Order for parents is important for some operations
	Parents []*Node

	// Order for children never matters
	Children []*Node

	Operation Operation

	inputs []int

	Output int
}

func NewNode(op Operation) *Node {
	return &Node{
		Operation: op,
		Children:  make([]*Node, 0, 3),
		inputs:    make([]int, 0, 3),
	}
}

func (n *Node) WireParents() {
	for _, p := range n.Parents {
		fmt.Println("adding self to parents children")
		p.Children = append(p.Children, n)
		p.WireParents()
	}
}

func (n *Node) Input(in int) {

	switch n.Operation {
	case OUTPUT:
		fmt.Println("saving output")
		n.Output = in
	case INPUT:
		fmt.Println("passing input to children")
		// pass it to each child
		for _, c := range n.Children {
			c.Input(in)
		}
	case PLUS:
		fmt.Println("saving input until input complete")
		n.inputs = append(n.inputs, in)
		if len(n.inputs) == 2 {
			fmt.Println("input completed, calculting output")
			output := n.inputs[0] + n.inputs[1]
			for _, c := range n.Children {
				c.Input(output)
			}
		}
	}
}

func main() {

	input1 := NewNode(INPUT)
	input2 := NewNode(INPUT)
	plus := NewNode(PLUS)
	plus.Parents = []*Node{input1, input2}
	output := NewNode(OUTPUT)
	output.Parents = []*Node{plus}

	output.WireParents()

	input1.Input(2)
	input1.Input(3)

	if output.Output != 5 {
		fmt.Println("failure ", output.Output)
	} else {
		fmt.Println("success")
	}

}
