#!/bin/python3

import math
import os
import random
import re
import sys
import itertools
from bisect import bisect_left

def busqueda(arr, l, tam, x):
    while l <= tam:
        medio = l + (tam-l) // 2
        if arr[medio] == x:
            return medio
        elif arr[medio] < x:
            tam = medio-1
        elif arr[medio] > x:
            l = medio+1
    if x > arr[medio]:
        medio-=1
    return medio

# Complete the climbingLeaderboard function below.
def climbingLeaderboard(scores, alice):
    scores2 = []
    resultados = [1]
    cuenta = 1
    for i in range(1, len(scores)):
        if scores[i] == scores[i-1]:
            resultados.append(cuenta)
        else:
            cuenta += 1
            resultados.append(cuenta)
    for a in alice:
        if scores[0] < a:
            scores2.append(1)
        elif scores[len(scores)-1] > a:
            scores2.append(resultados[len(resultados)-1]+1)
        else:
            aux = busqueda(scores, 0, len(scores), a)
            if scores[aux] == a:
                scores2.append(resultados[aux])
            else:
                scores2.append(resultados[aux]+1)
    return scores2


if __name__ == '__main__':
    scores_count = int(input())

    scores = list(map(int, input().rstrip().split()))

    alice_count = int(input())

    alice = list(map(int, input().rstrip().split()))

    result = climbingLeaderboard(scores, alice)

    print(result)
