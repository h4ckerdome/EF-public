package LinkedList

// Node Struct to hold values for Single Linked List
type Node struct {
	str  string
	next *Node
}

// Single Linked List
type LinkedList struct {
	head *Node
}

// Append function, should add new node to end of list
func (L *LinkedList) Append(str string) {
	n := Node{str: str}

	if L.head == nil {
		L.head = &n
		return
	}

	curr := L.head
	next := curr.next

	for next != nil {
		curr = next
		next = next.next
	}

	curr.next = &n
}

// Find a node with this string
func (L *LinkedList) Find(str string) *Node {
	// Insert Code Here
}

// Delete node, probably need to find the node first wink wink
func (L *LinkedList) Delete(n *Node) {
	// Insert Code Here
}

// Returns length of Linked List
func (L *LinkedList) Len() int {
	// Insert Code Here
}
