package main

import (
	"fmt"
	"goReStudy/tree"
)

type myTreeNode struct {

	node * tree.Node

}

func (myNode * myTreeNode) postOrder ()  {

	if myNode == nil || myNode.node == nil {

		return
	}

	left := myTreeNode{myNode.node.Left}
    right := myTreeNode{myNode.node.Right}
    left.postOrder()
	right.postOrder()
    myNode.node.Print()

}



func main()  {

	var  root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5,nil,nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateTreeNode(2)
	root.Print()
	fmt.Println()
	root.Travase()

	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()

}
