#!/bin/python3

import os
import sys


#
# Complete the bonetrousle function below.
#
def bonetrousle(n, k, b):
    resultado = []
    min = b*(b+1)//2
    max = k*b-b*(b-1)//2
    q = (n-min)//b
    r = (n-min)%b
    if min <= n <= max:
        if b == 1:
            resultado.append(n)
            return resultado
        for i in range(b):
            resultado.append(i+1+q)
        resultado[len(resultado)-r:] = [x+1 for x in resultado[len(resultado)-r:]]
    else:
        resultado.append(-1)
        return resultado
    return resultado


if __name__ == '__main__':

    t = int(input())

    for t_itr in range(t):
        nkb = input().split()

        n = int(nkb[0])

        k = int(nkb[1])

        b = int(nkb[2])

        result = bonetrousle(n, k, b)

        print(result)
