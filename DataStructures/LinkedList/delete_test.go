package LinkedList

import (
	"testing"
)

func TestDelete(t *testing.T) {
	// TODO: add head node
	head := Node{"One", nil}

	// TODO: create linked list
	ll := LinkedList{&head}

	ll.Append("Two")
	ll.Append("Three")
	// TODO: delete middle node

	curr := ll.head

	for curr != nil {
		if curr.str == "Two" {
			ll.Delete(curr)
		}
		curr = curr.next
	}

	// TODO: make sure middle node doesn't exist
	curr = ll.head

	found := false
	for curr != nil {
		if curr.str == "Two" {
			found = true
			break
		}
		curr = curr.next
	}
	if found {
		t.Fatalf(`ll.Delete("Two") failed to delete node "Two".`)
	}
}

func TestDeleteHead(t *testing.T) {
	node := Node{"One", nil}

	ll := LinkedList{&node}

	ll.Delete(&node)

	if ll.head != nil {
		t.Fatalf(`Failed to delete head %v`, ll.head.str)
	}
}
