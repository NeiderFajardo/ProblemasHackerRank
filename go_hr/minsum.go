package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
	"strings"
	"math"
	"sort"
)



/*
 * Complete the 'minSum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY num
 *  2. INTEGER k
 */

func minSum(num []int, k int32) int32 {
	// Write your code here
	suma := 0
	for i := 0; i <= int(k)-1; i++{
			sort.Ints(num)
			aux := math.Ceil(float64(num[len(num)-1])/float64(2))
			num[len(num)-1] = int(aux)
	}
	for _, i := range(num){
		suma += i
	}
	return int32(suma)
}


func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    numCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var num []int

    for i := 0; i < int(numCount); i++ {
        numItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        numItem := int32(numItemTemp)
        num = append(num, int(numItem))
    }

    kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := minSum(num, k)

    fmt.Println(result)
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