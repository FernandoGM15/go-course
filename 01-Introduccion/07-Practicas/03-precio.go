package main

import "fmt"

func getIGV(venta float32) float32 {
	return venta * 0.18
}

func main() {
	var venta float32

	fmt.Println("Ingrese el precio de venta del producto")
	fmt.Scanln(&venta)

	igv := getIGV(venta)
	fmt.Println("El IGV del producto es: ", igv)

	precioFinal := venta + igv
	fmt.Println("El precio final del producto es: ", precioFinal)
}
