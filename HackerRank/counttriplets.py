#!/bin/python3

import math
import os
import random
import re
import sys
from collections import Counter


def countTriplets(arr, r):
    triplet2 = Counter()
    triplet3 = Counter()
    resultado = 0
    for i in arr:
        if i in triplet3:
            resultado += triplet3[i]
        if i in triplet2:
            triplet3[i*r] += triplet2[i]
        triplet2[i*r] += 1
    return resultado

if __name__ == '__main__':

    nr = input().rstrip().split()

    n = int(nr[0])

    r = int(nr[1])

    arr = list(map(int, input().rstrip().split()))

    ans = countTriplets(arr, r)

    print(ans)