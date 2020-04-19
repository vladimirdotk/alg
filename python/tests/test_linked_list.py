import pytest
from typing import Iterator, Optional
from alg.linked_list import LinkedList, Node

@pytest.fixture
def nodes_generator() -> Iterator[Node]:
    return (Node(i) for i in range(1, 100000))

def test_append_head(nodes_generator: Iterator[Node]):
    n1 = next(nodes_generator)
    n2 = next(nodes_generator)
    
    ll = LinkedList(n1)    

    ll.append_head(n2)
    assert ll.head.value == n2.value

@pytest.mark.parametrize(
   ('head', 'tail', 'steps', 'expected'), [
       (Node(1), Node(2), 1, 2),
       (None, Node(3), 0, 3)
   ]
)
def test_append_tail(head: Optional[Node], tail: Node, steps: int, expected: int):
    ll = LinkedList(head)
    ll.append_tail(tail)
    
    current = ll.head
    while steps > 0:
        current = current.next
        steps -= 1
    
    assert current.value == expected


@pytest.mark.parametrize(
    ('nodes_cnt', 'position', 'expected'), [
        (0, 0, None),
        (2, 10, None),
        (1, 1 , Node(1).value),
        (2, 1 , Node(1).value),
        (10, 2, Node(2).value),
    ]
)
def test_get_position(nodes_generator: Iterator[Node], nodes_cnt: int, position: int, expected: Optional[Node]):
    if nodes_cnt < 1:
        return None

    ll = LinkedList(next(nodes_generator))
    while nodes_cnt != 0:
        nodes_cnt -= 1
        ll.append_tail(next(nodes_generator))    

    n = ll.get_position(position)

    if expected is None:
        assert n == None
    else:
        assert n.value == expected
