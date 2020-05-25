#!/bin/python3

import math
import os
import random
import re
import sys


#
# Complete the 'getStrength' function below.
#
# The function is expected to return an INTEGER.
# The function accepts following parameters:
#  1. STRING password
#  2. INTEGER weight_a
#

def getStrength(password, weight_a):
    # Complete the function
    aux = 97
    suma = 0
    lv = {}
    for i in range(weight_a, 26 + weight_a):
        lv[chr(aux)] = i % 26
        aux += 1
    print(lv)
    for j in password:
        suma += lv[j]
    return suma


if __name__ == '__main__':
    password = input()

    weight_a = int(input().strip())

    strength = getStrength(password, weight_a)

    print(strength)
