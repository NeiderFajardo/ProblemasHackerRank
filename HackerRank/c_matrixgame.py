#!/bin/python3


def definirganador(matriz, n,m):
    no_x = []
    no_y = []
    count = 1
    seguir = True
    ganador = 0
    nombres = ["Ashish", "Vivek"]
    for x in range(len(matriz)):
        for y in range(len(matriz[x])):
            if matriz[x][y] == 1:
                no_x.append(x)
                no_y.append(y)
    no_x = list(set(no_x))
    no_y = list(set(no_y))
    while (seguir):
        if len(no_x)+count > m or len(no_y)+count > n:
            seguir = False
        count+=1
        ganador+=1
    return nombres[ganador%2]





if __name__ == '__main__':
    test = int(input())
    for i in range(test):
        matriz = []
        n, m = map(int, input().split())
        for i in range(n):
            row = list(map(int, input().split()))
            matriz.append(row)
        result = definirganador(matriz, n, m)
        print(result)