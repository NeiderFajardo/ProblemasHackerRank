#!/bin/python3

import math
import os
import random
import re
import sys



#
# Complete the 'taskOfPairing' function below.
#
# The function is expected to return a LONG_INTEGER.
# The function accepts LONG_INTEGER_ARRAY freq as parameter.
#

def taskOfPairing(freq):
    # Write your code here
    res = 0
    sobra = []
    for i in range(len(freq)):
        if freq[i]%2 != 0 and 0 < i < len(freq)-1:
            if sobra[i-1] != 0:
                res += 1
                sobra.append(0)
            elif freq[i+1]%2 != 0:
                res += 1
                freq[i+1] -= 1
                sobra.append(0)
            else:
                sobra.append(1)
        else:
            if freq[i]%2 == 0:
                if 0 < i < len(freq)-1:
                    if sobra[i-1] != 0 and freq[i+1]%2 != 0:
                        freq[i] -= 2
                        res += 2
                        sobra.append(0)
                        sobra.append(0)
                    else:
                        sobra.append(0)
                else:
                    sobra.append(0)
            else:
                sobra.append(1)
        res += freq[i] // 2
    return res


if __name__ == '__main__':

    freq_count = int(input().strip())

    freq = []

    for _ in range(freq_count):
        freq_item = int(input().strip())
        freq.append(freq_item)

    result = taskOfPairing(freq)

    print(result)