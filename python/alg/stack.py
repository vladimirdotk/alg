from alg.linked_list import LinkedList, Node


class Stack():
    def __init__(self, linked_list: LinkedList):
        self.ll: LinkedList = linked_list
    
    def push(self, node: Node):
        self.ll.insert(node, 1)
    
    def pop(self):
        self.ll.delete(self.ll.head.value)