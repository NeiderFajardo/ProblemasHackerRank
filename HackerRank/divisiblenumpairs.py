#!/bin/python3


# Complete the divisibleSumPairs function below.
def divisibleSumPairs(n, k, ar):
    if len(ar) == 1:
        return 0
    aux = [x for x in ar[1:] if (ar[0]+x) % k == 0]
    return len(aux) + divisibleSumPairs(n, k, ar[1:])


if __name__ == '__main__':

    nk = input().split()

    n = int(nk[0])

    k = int(nk[1])

    ar = list(map(int, input().rstrip().split()))

    result = divisibleSumPairs(n, k, ar)

    print(result)
