package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList_AppendHead(t *testing.T) {
	tests := map[string]struct {
		ll         *LinkedList
		n          *Node
		expectedLl *LinkedList
	}{
		"append to empty list": {
			ll: &LinkedList{},
			n:  &Node{Value: 1},
			expectedLl: &LinkedList{
				Head: &Node{
					Value: 1,
				},
			},
		},
		"append to non empty list": {
			ll: &LinkedList{
				Head: &Node{
					Value: 1,
				},
			},
			n: &Node{
				Value: 2,
			},
			expectedLl: &LinkedList{
				Head: &Node{
					Value: 2,
					Next: &Node{
						Value: 1,
					},
				},
			},
		},
		"append empty node (will not change the linked list)": {
			ll: &LinkedList{
				Head: &Node{
					Value: 1,
				},
			},
			n: nil,
			expectedLl: &LinkedList{
				Head: &Node{
					Value: 1,
				},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tt.ll.AppendHead(tt.n)
			assert.Equal(t, tt.ll, tt.expectedLl)
		})
	}
}

func TestLinkedList_AppendTail(t *testing.T) {
	tests := map[string]struct {
		ll         *LinkedList
		n          *Node
		expectedLl *LinkedList
	}{
		"empty node": {
			ll: &LinkedList{
				Head: &Node{
					Value: 1,
				},
			},
			n: nil,
			expectedLl: &LinkedList{
				Head: &Node{
					Value: 1,
				},
			},
		},
		"empty linked list": {
			ll: &LinkedList{},
			n: &Node{
				Value: 1,
			},
			expectedLl: &LinkedList{
				Head: &Node{
					Value: 1,
				},
			},
		},
		"simple case": {
			ll: &LinkedList{
				Head: &Node{
					Value: 1,
					Next: &Node{
						Value: 2,
					},
				},
			},
			n: &Node{
				Value: 3,
			},
			expectedLl: &LinkedList{
				Head: &Node{
					Value: 1,
					Next: &Node{
						Value: 2,
						Next: &Node{
							Value: 3,
						},
					},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.ll.AppendTail(test.n)
			assert.Equal(t, test.ll, test.expectedLl)
		})
	}
}

func TestLinkedList_GetPosition(t *testing.T) {
	tests := map[string]struct {
		ll           *LinkedList
		pos          uint
		expectedNode *Node
	}{
		"zero position (not possible, we count from one), should return nil": {
			ll:           &LinkedList{},
			pos:          0,
			expectedNode: nil,
		},
		"returns head as a first node": {
			ll: &LinkedList{
				Head: &Node{
					Value: 1,
				},
			},
			pos: 1,
			expectedNode: &Node{
				Value: 1,
			},
		},
		"simple positive case": {
			ll: &LinkedList{
				Head: &Node{
					Value: 1,
					Next: &Node{
						Value: 2,
					},
				},
			},
			pos: 2,
			expectedNode: &Node{
				Value: 2,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			node := test.ll.GetPosition(test.pos)
			assert.Equal(t, node, test.expectedNode)
		})
	}
}

func TestLinkedList_Insert(t *testing.T) {
	tests := map[string]struct {
		ll         *LinkedList
		n          *Node
		pos        uint
		expectedLl *LinkedList
	}{
		"zero position not possible, does nothing": {
			ll: &LinkedList{},
			n: &Node{
				Value: 1,
			},
			pos:        0,
			expectedLl: &LinkedList{},
		},
		"insert at position 1": {
			ll: &LinkedList{
				Head: &Node{
					Value: 1,
				},
			},
			n: &Node{
				Value: 2,
			},
			pos: 1,
			expectedLl: &LinkedList{
				Head: &Node{
					Value: 2,
					Next: &Node{
						Value: 1,
					},
				},
			},
		},
		"no insertion at position 3, no such position": {
			ll: &LinkedList{
				Head: &Node{
					Value: 1,
				},
			},
			n: &Node{
				Value: 2,
			},
			pos: 3,
			expectedLl: &LinkedList{
				Head: &Node{
					Value: 1,
				},
			},
		},
		"insert at position 2": {
			ll: &LinkedList{
				Head: &Node{
					Value: 1,
					Next: &Node{
						Value: 2,
					},
				},
			},
			pos: 2,
			n: &Node{
				Value: 3,
			},
			expectedLl: &LinkedList{
				Head: &Node{
					Value: 1,
					Next: &Node{
						Value: 3,
						Next: &Node{
							Value: 2,
						},
					},
				},
			},
		},
		"insert at position 3 (as last node)": {
			ll: &LinkedList{
				Head: &Node{
					Value: 1,
					Next: &Node{
						Value: 2,
					},
				},
			},
			pos: 3,
			n: &Node{
				Value: 3,
			},
			expectedLl: &LinkedList{
				Head: &Node{
					Value: 1,
					Next: &Node{
						Value: 2,
						Next: &Node{
							Value: 3,
						},
					},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.ll.Insert(test.n, test.pos)
			assert.Equal(t, test.expectedLl, test.ll)
		})
	}
}
