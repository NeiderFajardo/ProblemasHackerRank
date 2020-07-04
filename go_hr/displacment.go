package main

import(
	"fmt"
	"strconv"
	"math"
)

func GenDisplaceFn(ac, v, d float64) func (float64) float64{
	res := func (ti float64) float64{
		s := ((0.5)*ac*math.Pow(ti,2)) + v*ti + d
		return s
	}
	return res
}

func main(){
	var acc, velocity, displacement float64
	aux := ""

	fmt.Print("Enter the value of acceleration: ")
	fmt.Scan(&aux)
	acc,_ = strconv.ParseFloat(aux, 64)
	fmt.Print("Enter the value of initial velocity: ")
	fmt.Scan(&aux)
	velocity,_ = strconv.ParseFloat(aux, 64)
	fmt.Print("Enter the value of initial displacement: ")
	fmt.Scan(&aux)
	displacement,_ = strconv.ParseFloat(aux, 64)

	Distancia := GenDisplaceFn(acc, velocity, displacement)

	fmt.Print("Enter the displacement time: ")
	fmt.Scan(&aux)
	si,_ := strconv.ParseFloat(aux, 64)

	fmt.Printf("Distance after %f seconds: %f", si, Distancia(si))
	fmt.Println()
}