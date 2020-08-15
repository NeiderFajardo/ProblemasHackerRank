package main

import (
	"math"	
	"fmt"
)

func factorial(n int) int{
	res := 1
	for i:=0;i<n;i++{
		res += res*i
	}

	return res
}

func main(){
	var t, n,m int
	p:= int(math.Pow(float64(10),float64(9))+7)
	
	fmt.Scan(&t)
	for i:=0; i<t;i++{
		fmt.Scan(&n,&m)
		m--
		num := factorial(n+m) 
		den:=factorial(n)*factorial(m)
		num=num/den
		fmt.Println(num%p)
	}
}