package main


import (
	"os"
	"fmt"
	"bufio"
)


func leerArchivo(nom string){
	data, err :=  os.Open(nom)
	if err != nil{
		fmt.Println("Error al leer el archivo.")
		return
	}
	fmt.Println("El archivo se leyo correctamente.")
	defer data.Close()
	
	apagarPermisos(nom)
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return
}

func apagarPermisos(nom string){
	err := os.Chmod(nom, 111)
	if err != nil{
		fmt.Println("Error al cambiar permisos.")
		return
	}
	fmt.Println("Se cambiaron los permisos sobre el archivo.")
}

func main(){
	nombre := "testfile.dms"
	leerArchivo(nombre)
}