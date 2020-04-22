from alg.binary_search import binary_search
from typing import List
import pytest

@pytest.mark.parametrize(
   ('sorted_list', 'value', 'index'), [
       ([1,2,3,4,5,6,7], 7, 6),
       ([1,2,3,4,5,6,7], 1, 0),
       ([1,2,3,4,5,6,7], 4, 3),
   ]
)
def test_binary_search(sorted_list: List, value: int, index: int):
    assert binary_search(sorted_list, value) == index