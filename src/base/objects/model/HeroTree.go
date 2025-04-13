package model

import "fmt"

type Hero struct {
	Id   int
	Name string
	left *Hero
	right *Hero
}

func PreOrder(node *Hero){
	if node != nil {
		PreOrder(node.left)
		fmt.Printf("节点信息: %d\n", node.Id)
		PreOrder(node.right)
	}
}