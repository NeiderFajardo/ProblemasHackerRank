#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the permutationEquation function below.
def permutationEquation(p):
    result = []
    for i in range(len(p)):
        pos = p.index(i+1)+1
        result.append(p.index(pos)+1)
    return result


if __name__ == '__main__':

    n = int(input())

    p = list(map(int, input().rstrip().split()))

    result = permutationEquation(p)

    print(result)