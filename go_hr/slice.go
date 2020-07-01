package main

import (
	"fmt"
	"strconv"
	"sort"
)

func esNumerico(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func main(){

		s := make([]int, 3)
		capacidad := 0
		p := ""
		for (p != "x"){
			fmt.Println("Please enter a number or x to exit:")
			fmt.Scan(&p)
			if esNumerico(p){
				num, _ :=  strconv.Atoi(p)
				if capacidad < 3 {
					for j:=0;j<3;j++{
						if s[j] == 0{
							s[j] = num
							break
						}
					}
					capacidad += 1
				}else{
					s = append(s, num)
				}
				sort.Ints(s)
			}
			for i, _ := range s{
				fmt.Print(s[i]," ")
			}
			fmt.Println("\n")
		}
}