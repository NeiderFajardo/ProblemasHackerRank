package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
	"strings"
)

/*
 * Complete the restaurant function below.
 */
func restaurant(l int, b int) int {
    /*
     * Write your code here.
	 */
	 var (
		 result, x int
		 es bool = true

	 )
	if l > b{
		x = l
	}else{
		x = b
	}
	for es{
		if (l%x == 0 && b%x == 0) || x == 1{
			es = false	 
		}else{
			x--
			fmt.Println(x)
		}
	}
	result = (l/x) * (b/x)
	return result
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        lb := strings.Split(readLine(reader), " ")

        lTemp, err := strconv.ParseInt(lb[0], 10, 64)
        checkError(err)
        l := int(lTemp)

        bTemp, err := strconv.ParseInt(lb[1], 10, 64)
        checkError(err)
        b := int(bTemp)

        result := restaurant(l, b)

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
