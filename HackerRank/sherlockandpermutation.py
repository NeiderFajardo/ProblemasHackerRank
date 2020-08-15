import math

q=int(input())

P=10**9+7

for _ in range(q):
    [n,m]=list(map(int,input().split()))
    m=m-1
    num=math.factorial(n+m)%P
    den=math.factorial(n)*math.factorial(m)
    den=pow(den,P-2, P)
    print(den)
    print((num*den)%P)