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

// Complete the appendAndDelete function below.
func appendAndDelete(s string, t string, k int32) string {
	aux := 0
	for i:=0; i < int(math.Min(float64(len(s)), float64(len(t)))); i++{
		if s[i] == t[i]{
			aux += 1
		}else{
			break;
		}
	}
	if (len(s)+len(t)-2*aux)>int(k){
		return "No"
	}else if (len(s)+len(t)-2*aux)%2 == int(k)%2{
		return "Yes"
	}else if(len(s)+len(t)-int(k)) < 0{
		return "Yes"
	} 
	return "No"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    s := readLine(reader)

    t := readLine(reader)

    kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := appendAndDelete(s, t, k)

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
