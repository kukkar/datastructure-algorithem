package linkedlist


type LL  struct{
	Head *Node
}

type Node struct {
	Value int
	Next *Node
}