import pytest
from typing import Optional
from alg.linked_list import LinkedList, Node

def test_append_head():
    n1 = Node(1)
    n2 = Node(2)
    
    ll = LinkedList(n1)    

    ll.append_head(n2)
    assert ll.head.value == 2

@pytest.mark.parametrize(
   ('head', 'tail', 'steps', 'expected'), [
       (Node(1), Node(2), 1, 2),
       (None, Node(3), 0, 3)
   ]
)
def test_append_tail(head: Optional[Node], tail: Node, steps: int, expected: Optional[int]):
    ll = LinkedList(head)
    ll.append_tail(tail)
    
    current = ll.head
    while steps > 0:
        current = current.next
        steps -= 1
    
    assert current.value == expected

def test_get_position():
    n1 = Node(1)
    n2 = Node(2)
    n3 = Node(3)

    ll = LinkedList(n1)
    ll.append_tail(n2)
    ll.append_tail(n3)

    n = ll.get_position(3)

    assert n.value == 3
