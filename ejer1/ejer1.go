package main

import(
	"fmt"
)

func main(){
	var nombre string
	var sueldo int
	fmt.Println("CUAL ES TU NOMBRE?")
	fmt.Scanln(&nombre)
	fmt.Println("CUAL ES TU SUELDO?")
	fmt.Scanln(&sueldo)
	var resultado int = sueldo + (sueldo / 10)
	fmt.Println(nombre, " tu suldo final es ", resultado)
}