#!/bin/python3


def getMoneySpent(keyboards, drives, b):
    maxi = 0
    if sorted(drives)[0] + sorted(keyboards)[0] > b:
        return -1
    for i in keyboards:
        for j in drives:
            if maxi < i+j <= b:
                maxi = i+j
    return maxi

if __name__ == '__main__':

    bnm = input().split()

    b = int(bnm[0])

    n = int(bnm[1])

    m = int(bnm[2])

    keyboards = list(map(int, input().rstrip().split()))

    drives = list(map(int, input().rstrip().split()))

    #
    # The maximum amount of money she can spend on a keyboard and USB drive, or -1 if she can't purchase both items
    #

    moneySpent = getMoneySpent(keyboards, drives, b)

    print(moneySpent)
