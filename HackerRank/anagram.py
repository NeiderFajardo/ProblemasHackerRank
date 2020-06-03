#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the anagram function below.
def anagram(s):
    if len(s)%2 != 0:
        return "-1"
    a = s[:len(s)//2]
    a = sorted(a)
    b = s[len(s) // 2:]
    b = sorted(b)
    for i in a:
        if i in b:
            b.remove(i)
    return len(b)


if __name__ == '__main__':

    q = int(input())

    for q_itr in range(q):
        s = input()

        result = anagram(s)

        print(result)


