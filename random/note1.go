package main

import(
	"fmt"
)

func main(){
	var num1,num2 int
	fmt.Println("Cuales son tus numeros? ")
	fmt.Scanln(&num1, &num2)
	resultado := num1 + num2
	fmt.Println("Tu resultado es ", resultado)
}