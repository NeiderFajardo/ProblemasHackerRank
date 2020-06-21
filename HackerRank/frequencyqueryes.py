#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the freqQuery function below.
from collections import defaultdict


def freqQuery(queries):

    x = defaultdict(set)
    y = defaultdict(int)
    for(q,n) in queries:
        print("entra")
        if q == 3: yield 1 if x[n] else 0; continue

        x[y[n]].discard(n)
        if q == 1 : y[n] += 1
        elif y[n] > 0: y[n] -=1
        x[y[n]].add(n)

if __name__ == '__main__':

    q = int(input().strip())

    queries = []

    for _ in range(q):
        queries.append(list(map(int, input().rstrip().split())))

    ans = freqQuery(queries)
    for i in ans:
        print(i)
