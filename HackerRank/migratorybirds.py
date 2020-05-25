#!/bin/python3


# Complete the migratoryBirds function below.
def migratoryBirds(arr):
    arr = sorted(arr)
    max = res = 0
    aux = 1
    for i in range(len(arr)-1):
        if arr[i] == arr[i+1]:
            aux += 1
        elif aux > max:
            max = aux
            aux = 1
            res = arr[i]
    if aux > max:
        res = arr[i]
    return res


if __name__ == '__main__':


    arr_count = int(input().strip())

    arr = list(map(int, input().rstrip().split()))

    result = migratoryBirds(arr)

    print(result)
