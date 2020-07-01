package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func findian(s string) string{
	s = strings.ToLower(s)
	if string(s[0]) == "i" && string(s[len(s)-1]) == "n" {
		if strings.Contains(s, "a"){
			return "Found"
		}
	}
	return "Not Found"
}

func main(){
	var palabra string
	fmt.Println("Please enter a string:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan(){
		palabra = scanner.Text()
	}
	res := findian(palabra)
	fmt.Println(res)
}