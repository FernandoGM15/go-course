package main

import "fmt"

func main() {
	/** App para restaurante
	* Descuentos de 10% hasta 100 de consumo
	* Descuentos de 20% hasta 200 de consumo
	* Descuentos de 30% mayor 200 de consumo
	* IGV 19%
	 */

	var consumo, descuento float64
	var datosDescuento string

	//  Entrada de datos
	fmt.Print("Ingrese consumo: ")
	fmt.Scanln(&consumo)

	if consumo >= 0 {
		if consumo <= 100 {
			//Descuento de 10%
			datosDescuento = "10%"
			descuento = consumo * 0.10
		} else if consumo > 100 && consumo <= 200 {
			//Descuento de 20%
			datosDescuento = "20%"
			descuento = consumo * 0.20
		} else if consumo > 200 {
			//Descuento de 30%
			datosDescuento = "30%"
			descuento = consumo * 0.30
		}
	} else {
		fmt.Println("Error al ingresar el consumo")
	}

	//Operaciones
	montoDescuento := consumo - descuento
	igv := montoDescuento * 0.19
	total := montoDescuento + igv
	fmt.Println("---------- Factura de consumo --------")
	fmt.Println("Descuento que se aplica: ", datosDescuento)
	fmt.Println("Consumo: ", consumo)
	fmt.Println("Descuento: ", descuento)
	fmt.Println("Monto con descuento: ", montoDescuento)
	fmt.Println("IGV: ", igv)
	fmt.Println("Total: ", total)
}
