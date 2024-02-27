package LinkedList

import (
	"testing"
)

func TestAppend(t *testing.T) {
	head := Node{"One", nil}

	ll := LinkedList{&head}

	ll.Append("Two")

	str := ll.head.next.str

	if str != "Two" {
		t.Fatalf(`Failed to add new node with value "Two". str: %v`, str)
	}
}

func TestAppendEmptyList(t *testing.T) {
	ll := LinkedList{}

	ll.Append("One")

	str := ll.head.str

	if str != "One" {
		t.Fatalf(`Failed to add new node with value "One" to empty list. str: %v`, str)
	}
}
