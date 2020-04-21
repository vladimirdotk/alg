"""Make a Queue class using a list!
Hint: You can use any Python list method
you'd like! Try to write each one in as 
few lines as possible.
Make sure you pass the test cases too!"""

from alg.linked_list import Node
from typing import Optional

class Queue:
    def __init__(self, head: Optional[Node] = None):
        self.storage = [head]

    def enqueue(self, new_element: Node):
        self.storage.append(new_element)

    def peek(self) -> Optional[Node]:
        if self.storage:
            return self.storage.pop(0)
        return None

    def dequeue(self):
        if self.storage:
            self.storage.pop()