#!/bin/python3

import os
import sys

# Complete the solve function below.
def solve(n):
    res = 0
    nu = 1
    aux = 0
    while res <= n:
        if res == n:
            return "IsFibo"
        aux = res
        res += nu
        nu = aux
    return "IsNotFibo"

if __name__ == '__main__':
    t = int(input())

    for t_itr in range(t):
        n = int(input())

        result = solve(n)

        print(result)
