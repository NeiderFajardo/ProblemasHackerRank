package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"strconv"
)

func Swap(nums []int, pos int){
	aux := nums[pos]
	nums[pos] = nums[pos+1]
	nums[pos+1] = aux
}

func BubbleSort(nums []int){
	for i:=0; i<len(nums)-1;i++{
		for j:=0;j<len(nums)-i-1;j++{
			if nums[j] > nums[j+1]{
				Swap(nums, j)
			}
		}
	}
}


func main(){
	var numeros []int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter the numbers in one line, separate with single spaces:")
	if scanner.Scan(){
		linea := strings.Split(scanner.Text(), " ")
		for _, num := range linea{
			if aux, err := strconv.Atoi(num); err == nil{
				numeros = append(numeros, aux)
			}
		}
		BubbleSort(numeros)
	}
	for _, n := range numeros{
		fmt.Print(strconv.Itoa(n)," ")
	}
	fmt.Println()
}