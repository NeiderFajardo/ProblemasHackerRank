#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the arrayManipulation function below.
def arrayManipulation(n, queries):
    count = 0
    for i in range(1,n+1):
        queries[i] += queries[i-1]
        count = max(queries[i], count)
    return count


if __name__ == '__main__':

    nm = input().split()

    n = int(nm[0])

    m = int(nm[1])

    queries = []
    arr = [0]*(n+2)
    for i in range(m):
        queries.append(list(map(int, input().rstrip().split())))
        arr[queries[i][0]] += queries[i][2]
        arr[queries[i][1]+1] -= queries[i][2]
    result = arrayManipulation(n, arr)

    print(result)
