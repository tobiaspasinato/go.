package main

import(
	"fmt"
)

func main(){
	var contador int = 0
	var contador_pares int = 0
	var numero, numero_menor, numero_mayor int
	var bandera_menor bool = true
	var bandera_mayor bool = true
	for true{
		fmt.Println("Ingrese un numero")
		fmt.Scanln(&numero)
		if numero % 2 == 0{
			if bandera_mayor == true || numero_mayor < numero{
				numero_mayor = numero
				bandera_mayor = false
			}
			contador_pares += 1
		}
		if bandera_menor == true || numero_menor > numero{
			numero_menor = numero
			bandera_menor = false
		}
		if contador < 5{
			break
		}
		contador += 1
	}
}