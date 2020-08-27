package main

import (
	"fmt"
)

func modP(a, m int) int {
	var res int = a % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

func mult(p0, p1, x0, x1, m int) (int, int) {
	return modP((p0*x0 - p1*x1), m), modP((p0*x1 + p1*x0), m)
}

func solve(a, b, k, m int) (int, int) {
	x0, x1 := a, b
	p0, p1 := 1, 0
	for k > 0 {
		if k%2 != 0 {
			p0, p1 = mult(p0, p1, x0, x1, m)
		}
		k = k / 2
		x0, x1 = mult(x0, x1, x0, x1, m)
	}
	return p0, p1
}

func main() {
	var q, a, b, k, m int
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		fmt.Scan(&a, &b, &k, &m)
		c, d := solve(a, b, k, m)
		fmt.Printf("%v %v\n", c, d)
	}
}
