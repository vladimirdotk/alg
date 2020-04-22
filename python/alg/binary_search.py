from math import floor
from typing import List

def binary_search(sorted_list: List,  value: int) -> int:
    left = 0
    right = len(sorted_list) - 1
    while left <= right:
        mid = (left + right)//2

        if sorted_list[mid] == value:
            return mid
        
        if sorted_list[mid] < value:
            left = mid + 1
        else:
            right = mid -1
    
    return -1

    

if __name__ == '__main__':
    print(binary_search([1,2,3,4,5,6,7], 1))