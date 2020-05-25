#!/bin/python3


# Complete the birthday function below.
def birthday(s, d, m):
    pos = res = 0
    while (pos + m <= len(s)):
        aux = s[pos:pos+m]
        if sum(aux) == d:
            res += 1
        pos += 1
    return res


if __name__ == '__main__':

    n = int(input().strip())

    s = list(map(int, input().rstrip().split()))

    dm = input().rstrip().split()

    d = int(dm[0])

    m = int(dm[1])

    result = birthday(s, d, m)

    print(result)
