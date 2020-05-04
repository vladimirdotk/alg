package structures

// Node is a part of Linked List
type Node struct {
	Value int
	Next  *Node
}

// NewNode is a Node constructor
func (n *Node) NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

// LinkedList is a linked list data structure
type LinkedList struct {
	Head *Node
}

// LinkedList is a linked list constructor
func (ll *LinkedList) NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// AppendHead appends a node to the head of a linked list
// If node is not nil, otherwise do nothing
func (ll *LinkedList) AppendHead(node *Node) {
	if node == nil {
		return
	}

	if ll.Head != nil {
		current := ll.Head
		ll.Head = node
		ll.Head.Next = current
	}

	ll.Head = node
}
