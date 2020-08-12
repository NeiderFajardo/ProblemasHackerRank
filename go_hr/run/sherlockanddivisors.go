package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
	"strings"
	"math"
)


/*
 * Complete the divisors function below.
 */
func divisors(n int) int {
    /*
     * Write your code here.
	 */
	res := 0
	x := math.Sqrt(float64(n))
	for i:=1;i<int(x)+1;i++{
		if n%i==0 && i%2==0{
			res++
		}
		if n%(n/i)==0 && (n/i)%2==0{
			res++
		}
		if i ==n/i && i%2 == 0 && n%i == 0{
			res--
		}
	}
	return res
}

func main() {

    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int(tTemp)
	
    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        n := int(nTemp)

		result := divisors(n)
		
        fmt.Println(result)
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
