#!/bin/python3


# Complete the pageCount function below.
def pageCount(n, p):
    if int((n-p)/2) < int(p/2):
        if n%2==0 and p%2!=0:
            return int((n-p)/2)+1
        return int((n-p)/2)
    return int(p / 2)


if __name__ == '__main__':

    n = int(input())

    p = int(input())

    result = pageCount(n, p)

    print(result)
