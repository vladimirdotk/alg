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
