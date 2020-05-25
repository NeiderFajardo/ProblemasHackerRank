#!/bin/python3

import os

# Complete the hourglassSum function below.
def hourglassSum(arr):
    sumas = []
    for i in range(4):
        for j in range(4):
            sumas.append(arr[i][j]+arr[i][j+1]+arr[i][j+2]+arr[i+1][j+1]+arr[i+2][j]+arr[i+2][j+1]+arr[i+2][j+2])
    return max(sumas)


if __name__ == '__main__':
    arr = []

    for _ in range(6):
        arr.append(list(map(int, input().rstrip().split())))

    result = hourglassSum(arr)

    print(result)
