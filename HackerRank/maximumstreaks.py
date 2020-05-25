#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'getMaxStreaks' function below.
#
# The function is expected to return an INTEGER_ARRAY.
# The function accepts STRING_ARRAY toss as parameter.
#

def getMaxStreaks(toss):
    # Return an array of two integers containing the maximum streak of heads and tails respectively
    head = 0
    tail = 0
    maxim_head = [0]
    maxim_tail = [0]
    anterior = toss[0]
    for i in toss:
        if anterior != i:
            maxim_head.append(head)
            maxim_tail.append(tail)
            head = 0
            tail = 0
        if i == "Heads":
            head += 1
        else:
            tail += 1
        anterior = i
    resultado = [max(maxim_head), max(maxim_tail)]
    return resultado


if __name__ == '__main__':

    toss_count = int(input().strip())

    toss = []

    for _ in range(toss_count):
        toss_item = input()
        toss.append(toss_item)

    ans = getMaxStreaks(toss)

    print(ans)