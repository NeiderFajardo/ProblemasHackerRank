#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the extraLongFactorials function below.
def extraLongFactorials(n):
    resultado = n
    for i in range(n//2):
        if (n-i-1) != i+1:
            resultado *= (n-i-1)*(i+1)
        else:
            resultado *= i+1
    print(resultado)


if __name__ == '__main__':
    n = int(input())

    extraLongFactorials(n)
