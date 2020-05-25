#!/bin/python3

# Complete the rotLeft function below.
def rotLeft(a, d):
    aux = []
    aux = a.copy()
    for i in range(len(a)):
        if i - d < 0:
            aux[len(a) + i - d] = a[i]
        else:
            aux[i - d] = a[i]
    return aux


if __name__ == '__main__':

    nd = input().split()

    n = int(nd[0])

    d = int(nd[1])

    a = list(map(int, input().rstrip().split()))

    result = rotLeft(a, d)

    print(result)