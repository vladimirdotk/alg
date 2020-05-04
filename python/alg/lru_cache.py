from typing import Dict, List
import random

def lru_cache(func):
    cache = {}
    def inner(*args: List, **kwargs: Dict):
        nonlocal cache
        key = str(args) + str(kwargs)
        if key in cache:
            print("Will use cache")
            return cache[key]
        
        result = func(*args, **kwargs)
        cache[str(args) + str(kwargs)] = result
        return result

    return inner

@lru_cache
def my_func(x: int, y: int) -> int:
    return random.randint(x,y)

if __name__ == '__main__':
    print(my_func(100,200))
    print(my_func(100,200))
    print(my_func(100,200))
    print(my_func(1,301))