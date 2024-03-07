package main

import (
	"fmt"
	"reflect"


)

func main(){
	fmt.Println("Hello Go")
	mensaje := "El resultado es "
	var num1 int = 4 //no tengo que darle el tipo de dato
	var num2 int //se puede declarar despues el valor
	num2 = 5
	var resultado int = num1 + num2
	fmt.Println(mensaje, resultado)
	fmt.Println(reflect.TypeOf(mensaje))
}