#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the makingAnagrams function below.
def makingAnagrams(s1, s2):
    s1 = sorted(s1)
    s2 = sorted(s2)
    resp = 0
    for i in s1:
        if i in s2:
            s2.remove(i)
        else:
            resp += 1
    return resp+len(s2)

if __name__ == '__main__':

    s1 = input()

    s2 = input()

    result = makingAnagrams(s1, s2)

    print (result)