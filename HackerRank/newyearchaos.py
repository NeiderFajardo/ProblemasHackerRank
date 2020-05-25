#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the minimumBribes function below.
def minimumBribes(q):
    aux = []
    c = 0
    for i in range(len(q)):
        aux.append(i+1)
    for j in range(len(q)-1):
        if q[j] != aux[j]:
            if q[j] != aux[j+1] and q[j] - (j+1) <= 2:
                n = aux[j]
                aux[j] = aux[j+2]
                aux[j+2] = aux[j+1]
                aux[j+1] = n
                c += 2
            elif q[j] - (j+1) > 2:
                return "Too chaotic"
            else:
                n = aux[j]
                aux[j] = aux[j+1]
                aux[j+1] = n
                c += 1
    if aux != q:
        return "Too chaotic"
    return c


if __name__ == '__main__':
    t = int(input())

    for t_itr in range(t):
        n = int(input())

        q = list(map(int, input().rstrip().split()))

        result = minimumBribes(q)
        print(result)