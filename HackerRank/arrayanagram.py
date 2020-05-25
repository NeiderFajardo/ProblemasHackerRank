#!/bin/python3

import math
import os
import random
import re
import sys

def es_anagrama(x, y):
    x = sorted(x)
    y = sorted(y)
    if x == y:
        return True
    return False

# Complete the sherlockAndAnagrams function below.
def sherlockAndAnagrams(s):
    result = 0
    for i in range(len(s)-1):
        n = 1
        aux = 1
        while n +i < len(s):
            for j in range(len(s)-n-i):
                if es_anagrama(s[i:i+n], s[i+aux+j:i+n+aux+j]):
                    result += 1
            n += 1
        aux += 1
    return result


if __name__ == '__main__':

    q = int(input())

    for q_itr in range(q):
        s = input()

        result = sherlockAndAnagrams(s)

        print(result,"\n")

