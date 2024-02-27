package LinkedList

import (
	"testing"
)

func TestFind(t *testing.T) {
	head := Node{"One", nil}

	ll := LinkedList{&head}

	ll.Append("Two")
	ll.Append("Three")
	ll.Append("Four")

	node := ll.Find("Three")

	if node.str != "Three" {
		t.Fatalf(`Failed to find now with "Three" as value. Node.str = %v`, node.str)
	}
}
