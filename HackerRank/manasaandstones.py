#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the stones function below.
def stones(n, a, b):
    res = []
    a_aux = 0
    for i in range(n):
        suma = 0
        b_aux = n-1-a_aux
        for j in range(a_aux):
            suma += a
        for k in range(b_aux):
            suma += b
        a_aux += 1
        res.append(suma)
    res = list(set(res[::-1]))
    return sorted(res)


if __name__ == '__main__':

    T = int(input())

    for T_itr in range(T):
        n = int(input())

        a = int(input())

        b = int(input())

        result = stones(n, a, b)

        print(result)

