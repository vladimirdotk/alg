from alg.queue import Queue, Node
import pytest

def test_enqueue():
    q = Queue()
    q.enqueue(Node(10))
    assert q.storage.pop(-1).value == 10

def test_peek():
    q = Queue()
    assert q.peek() is None

    q.storage.append(Node(10))
    q.storage.append(Node(100))
    assert q.peek().value == 10

def test_dequeue():
    q = Queue()
    q.dequeue()

    q.storage.append(Node(100))
    q.storage.append(Node(200))
    q.dequeue()

    assert q.storage[0].value == 100

