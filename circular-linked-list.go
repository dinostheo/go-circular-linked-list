package cqlnls

type Node struct {
	Value          interface{}
	Next, Previous *Node
}

type CircularLinkedList struct {
	head, tail *Node
	len        int
}

/*
 - AddPrevious() *Node
 - AddNext() *Node
 - Previous() *Node
 - Next() *Node
 - Remove() *Node
 - Length() int
*/
