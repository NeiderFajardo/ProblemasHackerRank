#!/bin/python3


def maxad(lista):
    maximum = float('-inf')
    for element, adjacent in zip(lista[:-1], lista[1:]):
        candidate = element *adjacent
        if candidate > maximum:
            maximum =candidate
    return maximum

d = [3,6,-2,-5,7,3]
c = [1]
print(maxad(c))