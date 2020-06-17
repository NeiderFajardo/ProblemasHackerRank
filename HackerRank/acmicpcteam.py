#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the acmTeam function below.
def acmTeam(topic):
    res = []
    m = len(topic)
    x = 1
    maximo = 0
    while x < m:
        for i in topic[x:]:
            cuenta = 0
            for j in range(len(i)):
                if i[j] == '1' or topic[x-1][j] == '1':
                    cuenta += 1
            if cuenta > maximo:
                maximo = cuenta
                res = []
                res.append(1)
            elif cuenta == maximo:
                res.append(1)
        x += 1
    return [maximo, len(res)]


if __name__ == '__main__':

    nm = input().split()

    n = int(nm[0])

    m = int(nm[1])

    topic = []

    for _ in range(n):
        topic_item = input()
        topic.append(topic_item)

    result = acmTeam(topic)

    print(result)
