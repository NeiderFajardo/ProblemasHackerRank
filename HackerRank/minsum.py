#!/bin/python3
import math



def minsum(num, k):
    num = sorted(num)
    num = num[::-1]
    print(num)
    for i in range(k):
        aux = num.index(max(num))
        num[aux] = math.ceil(num[aux]/2)
    print(num)
    return sum(num)



if __name__ == '__main__':
    tam = int(input())
    num = []
    for i in range(tam):
        num.append(int(input()))
    k = int(input())
    result = minsum(num, k)
    print(result)