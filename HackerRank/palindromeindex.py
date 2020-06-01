#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the palindromeIndex function below.
def palindromeIndex(s):
    result = 0
    for i in range(len(s)//2):
        if s[i] != s[len(s)-1-i]:
            aux = s[0:i]+s[i+1:]
            if aux == aux[::-1]:
                return i
            return len(s)-1-i
    return "-1"

if __name__ == '__main__':

    q = int(input())

    for q_itr in range(q):
        s = input()

        result = palindromeIndex(s)

        print(result)