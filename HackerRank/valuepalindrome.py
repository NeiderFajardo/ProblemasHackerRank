#!/bin/python3

import math
import os
import random
import re
import sys


# Complete the highestValuePalindrome function below.
def highestValuePalindrome(s, n, k):
    if len(s) == 1 and k >= 1:
        return "9"
    elif len(s) == 1 and k < 1:
        return s
    lista = list(s)
    left = lista[:len(lista) // 2]
    m = ""
    if len(lista) % 2 == 0:
        right = lista[len(lista) // 2:]
        right = right[::-1]
    else:
        right = lista[1 + (len(lista) // 2):]
        right = right[::-1]
        m = s[len(s) // 2:(len(s) // 2) + 1]
    for i in range(len(left)):
        if left[i] != right[i]:
            if k >= 2:
                left[i] = "9"
                right[i] = "9"
                k -= 2
            elif k > 0:
                if int(left[i]) > int(right[i]):
                    right[i] = left[i]
                else:
                    left[i] = right[i]
                k -= 1
        elif left[i] == right[i] and left[i] != "9" and k >= 2:
            left[i] = "9"
            right[i] = "9"
            k -= 2
    respuesta = "".join(left)
    if k > 0 and m != "":
        result = respuesta + "9" + "".join(right[::-1])
    else:
        result = respuesta + m + "".join(right[::-1])
    if result == result[::-1]:
        return result
    return "-1"


if __name__ == '__main__':

    nk = input().split()

    n = int(nk[0])

    k = int(nk[1])

    s = input()

    result = highestValuePalindrome(s, n, k)

    print(result)
