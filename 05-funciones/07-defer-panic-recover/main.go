package main

import (
	"fmt"
	"os"
)

func main() {

	//Function for handling panics at the end of the program
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Ooops, al parecer el programa no finalizo de forma correcta")
		}
	}()

	if file, error := os.Open("05-funciones/07-defer-panic-recover/heallo.txt"); error != nil {
		// Throw panic to force the programs closing
		panic("No fue posible leer el archivo")
	} else {
		defer func() {
			fmt.Println("Closing the file")
			file.Close()
		}()
		content := make([]byte, 254)
		long, _ := file.Read(content)
		fileContent := string(content[:long])
		fmt.Println(fileContent)
	}
}
