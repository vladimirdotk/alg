from typing import Optional

class Node:
    def __init__(self, value: int):
        self.value = value
        self.next: Optional[Node] = None

class LinkedList:
    def __init__(self, node: Node = None):
        self.head = node

    def append_head(self, node: Node):
        current = self.head
        self.head = node
        self.head.next = current
    
    def append_tail(self, node: Node):
        if self.head is None:
            self.head = node
            return
       
        current = self.head
        while current.next:
            current = current.next
        current.next = node

    def get_position(self, position: int) -> Optional[Node]:
        if position < 1:
            return None
        
        current = self.head
        while current:
            position -= 1
            if position == 0:
                return current
            current = current.next
        
        return None

    def insert(self, node: Node, position: int):
        if position < 1:
            return
        
        current = self.head
        if position == 1:
            self.head = node
            self.head.next = current
            return
        
        while current:
            position -= 1
            if position == 1:
                old_next = current.next
                current.next = node
                node.next = old_next
                return
            current = current.next