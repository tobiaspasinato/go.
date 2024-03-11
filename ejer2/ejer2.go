package main

import(
	"fmt"
)

func main(){
	var edad int
	fmt.Scanln(&edad)
	if edad >= 18 {
		fmt.Println("Sos mayor de edad")
	} else {
		if edad <= 12 {
			fmt.Println("sos un infante")
		} else {
			fmt.Println("sos un adolecente")
		}
	}
}