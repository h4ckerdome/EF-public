package LinkedList

import (
	"testing"
)

func TestLen(t *testing.T) {
	head := Node{"One", nil}

	ll := LinkedList{&head}

	ll.Append("Two")
	ll.Append("Three")
	ll.Append("Four")

	len := ll.Len()

	if len != 4 {
		t.Fatalf(`Failed to return 4 for length, len = %v`, len)
	}
}
