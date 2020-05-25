#!/bin/python3

import math
import os
import random
import re
import sys

def contiene(x, y):
    i = 0
    r = 0
    while i <= len(x)-1:
        j = 0
        while j <= len(y)-1 and i <= len(x)-1:
            if x[i] == y[j]:
                i += 1
                y = y[j+1:]
                j = 0
                r += 1
            else:
                j += 1
        i += 1
    return r
# Complete the commonChild function below.
def commonChild(s1, s2):
    result = 0
    for i in range(len(s1)):
        aux = contiene(s1[i:], s2)
        if aux > result:
            result = aux
    return result


if __name__ == '__main__':

    s1 = input()

    s2 = input()

    result = commonChild(s1, s2)

    print(result)
