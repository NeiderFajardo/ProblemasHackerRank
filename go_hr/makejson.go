package main

import(
	"fmt"
	"encoding/json"
	"bufio"
	"os"
)

func mostrarjson(m map[string]string){
	if pjson, err := json.MarshalIndent(m, "", "	"); err == nil{
		fmt.Println(string(pjson))
	}
}

func main(){
	p := make(map[string]string)
	var name, addr string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter a name: ")
	if scanner.Scan(){
		name = scanner.Text()
	}
	fmt.Print("Please enter an address: ")
	if scanner.Scan(){
		addr = scanner.Text()
	}
	p["address"] = addr
	p["name"] = name
	mostrarjson(p)

}