package tree

import "fmt"

type Node struct {

	 Value int
	 Left ,Right  *Node

}

func (node Node) Print ()  {
	fmt.Print(node.Value," ")
}

func (node *Node) Travase()  {

	if node == nil {
		return
	}
	node.Left.Travase()
	node.Print()
	node.Right.Travase()

}

func  CreateTreeNode(value int) (tree  *Node)  {

	return  &Node{Value: value}
}


