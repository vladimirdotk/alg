from alg.stack import Stack
from alg.linked_list import LinkedList, Node


def test_stack():
    stack = Stack(LinkedList(Node(1)))
    stack.push(Node(2))
    assert stack.ll.head.value == 2

    stack.pop()
    assert stack.ll.head.value == 1
