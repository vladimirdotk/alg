import pytest
from typing import Iterator, Optional
from alg.linked_list import LinkedList, Node

def node_generator() -> Iterator[Node]:
    return (Node(i) for i in range(1, 100000))

def get_linked_list(size: int) -> LinkedList:
    gen = node_generator()
    
    ll = LinkedList(next(gen))
    while size != 0:
        size -= 1
        ll.append_tail(next(gen)) 

    return ll

def test_append_head():
    gen = node_generator()

    n1 = next(gen)
    n2 = next(gen)
    
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
def test_get_position(nodes_cnt: int, position: int, expected: Optional[Node]):
    if nodes_cnt < 1:
        return None

    ll = get_linked_list(nodes_cnt)  
    n = ll.get_position(position)

    if expected is None:
        assert n == None
    else:
        assert n.value == expected

@pytest.mark.parametrize(
    ('ll_size', 'node', 'pos', 'expected'), [
        (5, Node(10), 1, 10),
        (5, Node(10), 2, 10),
        (5, Node(10), 5, 10),
        (1, Node(1), 1, 1)
    ]
)
def test_insert(ll_size: int, node: Node, pos: int, expected: int):
    ll = get_linked_list(ll_size)
    ll.insert(node, pos)

    if pos == 1:
        assert ll.head.value == expected
        return

    current = ll.head
    while current:
        pos -= 1
        if pos == 0:
            assert current.value == expected
            return
        current = current.next
    
    raise Exception("Failed to test insertion")