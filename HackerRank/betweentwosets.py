#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'getTotalX' function below.
#
# The function is expected to return an INTEGER.
# The function accepts following parameters:
#  1. INTEGER_ARRAY a
#  2. INTEGER_ARRAY b


def getTotalX(a, b):
    result = []
    for i in range(a[len(a)-1], b[0]+1):
        a_aux = [i % x for x in a]
        b_aux = [x % i for x in b]
        if sum(a_aux) + sum(b_aux) == 0:
            result.append(i)
    return len(result)

if __name__ == '__main__':
    first_multiple_input = input().rstrip().split()

    n = int(first_multiple_input[0])

    m = int(first_multiple_input[1])

    arr = list(map(int, input().rstrip().split()))

    brr = list(map(int, input().rstrip().split()))

    total = getTotalX(arr, brr)

    print(total)
