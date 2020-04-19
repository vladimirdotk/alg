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

    def get_position(self, position) -> Optional[Node]:
        if position < 1:
            return None
        
        if position == 1:
            return self.head
        
        current = self.head
        while current:
            position -= 1
            if position == 0:
                return current
            current = current.next
        
        return None