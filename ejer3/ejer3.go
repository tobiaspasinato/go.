package main

import(
	"fmt"
)

func main(){
	var contador int = 0
	var contador_pares int = 0
	var contador_impares int = 0
	var acumulador_positivos int = 0
	var acumulador_negativos int = 0
	var numero, numero_menor, numero_mayor int
	var bandera_menor bool = true
	var bandera_mayor bool = true
	for contador < 5{
		fmt.Println("Ingrese un numero")
		fmt.Scanln(&numero)
		if numero % 2 == 0{
			if bandera_mayor == true || numero_mayor < numero{
				numero_mayor = numero
				bandera_mayor = false
			}
			contador_pares += 1
		} else {
			contador_impares += 1
		}
		if bandera_menor == true || numero_menor > numero{
			numero_menor = numero
			bandera_menor = false
		}
		if numero > 0{
			acumulador_positivos += numero
		} else {
			acumulador_negativos += numero
		}
		contador += 1
	}
	fmt.Println("contador de pares ", contador_pares, " contador de impares ", contador_pares)
	fmt.Println("numero menor ingresado ", numero_menor)
	fmt.Println("de los pares el numero mayor ", numero_mayor)
	fmt.Println("suma de los positivos ", acumulador_positivos)
	fmt.Println("Producto de los negativos", (acumulador_negativos * contador_impares))

}