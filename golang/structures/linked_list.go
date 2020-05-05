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

func (ll *LinkedList) AppendTail(node *Node) {
	if node == nil {
		return
	}

	if ll.Head == nil {
		ll.Head = node
		return
	}

	current := ll.Head

	for {
		if current.Next == nil {
			break
		}
		current = current.Next
	}

	current.Next = node
}

func (ll *LinkedList) GetPosition(pos uint) *Node {
	if pos == 0 {
		return nil
	}

	current := ll.Head

	for {
		if current == nil {
			return nil
		}
		pos--
		if pos == 0 {
			return current
		}
		current = current.Next
	}
}

func (ll *LinkedList) Insert(node *Node, pos uint) {
	if pos == 0 {
		return
	}

	if pos == 1 {
		tmp := ll.Head
		ll.Head = node
		node.Next = tmp
		return
	}

	prevLink := ll.Head

	for {
		if prevLink == nil {
			break
		}
		pos--
		if pos == 1 {
			tmp := prevLink.Next
			prevLink.Next = node
			node.Next = tmp
			break
		}
		prevLink = prevLink.Next
	}
}

func (ll *LinkedList) Delete(value int) {
	if ll.Head == nil {
		return
	}

	if ll.Head.Value == value {
		ll.Head = ll.Head.Next
	}

	current := ll.Head
	for {
		previous := current
		if previous == nil {
			break
		}
		current = current.Next
		if current == nil {
			break
		}
		if current.Value == value {
			previous.Next = current.Next
			break
		}
	}
}
