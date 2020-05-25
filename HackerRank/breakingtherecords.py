
# Complete the breakingRecords function below.
def breakingRecords(scores):
    min = max = scores[0]
    aux_min = aux_max = 0
    for i in scores:
        if i < min:
            min = i
            aux_min+=1
        elif i > max:
            max = i
            aux_max+=1
    return [aux_max, aux_min]


if __name__ == '__main__':

    n = int(input())

    scores = list(map(int, input().rstrip().split()))

    result = breakingRecords(scores)

    print(result)
