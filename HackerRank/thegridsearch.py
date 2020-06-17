#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the gridSearch function below.
def gridSearch(G, P):
    n = len(P)
    m = len(P[0])
    for i in range(len(G)-m):
        for k in range(len(G[0])-n):
            count = 0
            for j in range(n):
                for l in range(m):
                    print(G[i+j][k+l], " ", P[j][l])
                    if G[i+j][k+l] == P[j][l]:
                        count += 1
                        if count == n*m:
                            return "YES"
    return "NO"



if __name__ == '__main__':

    t = int(input())

    for t_itr in range(t):
        RC = input().split()

        R = int(RC[0])

        C = int(RC[1])

        G = []

        for _ in range(R):
            G_item = input().split()
            G.append(G_item)

        rc = input().split()

        r = int(rc[0])

        c = int(rc[1])

        P = []

        for _ in range(r):
            P_item = input().split()
            P.append(P_item)

        result = gridSearch(G, P)

        print(result)