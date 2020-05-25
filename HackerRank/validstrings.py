#!/bin/python3

import math
import os
import random
import re
import sys
from itertools import chain
from collections import Counter


# Complete the isValid function below.
def isValid(s):
    aux = {}
    hay = False
    for i in s:
        if i in aux:
            aux[i] += 1
        else:
            aux[i] = 1
    if len(aux) == 1:
        return "YES"
    for key, val in aux.items():
        if 1 == val:
            aux.pop(key)
            hay = True
            break
    min_aux = min(aux.keys(), key=(lambda k: aux[k]))
    max_aux = max(aux.keys(), key=(lambda k: aux[k]))
    c_min = [x for x in aux.values()].count(aux[min_aux])
    c = [j for j in aux.values()].count(aux[max_aux])
    if c != c_min and hay:
        return "NO"
    if (aux[max_aux] - 1 == aux[min_aux] and hay is False) or max_aux == min_aux:
        return "YES"
    return "NO"


if __name__ == '__main__':
    s = input()

    result = isValid(s)

    print(result)
