package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func reversegame(n,k int) int {
	aux := 0
	nums := make(map[int]int)
	for i := 0; i < n/2; i++{
		nums[n-i-1] = aux 
		aux++
		nums[i] = aux
		aux++
	}
	if n%2 != 0 {
		nums[n/2] = n-1
	}
	fmt.Println(nums)
	return nums[k]
}


func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)
	
	for i := 0; i < int(t); i++{
		nk := strings.Split(readLine(reader), " ")

    	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
    	checkError(err)
		n := int(nTemp)

		kTemp, err := strconv.ParseInt(nk[1], 10, 64)
		checkError(err)
		k := int(kTemp)

		fmt.Println(reversegame(n,k))
	} 
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
