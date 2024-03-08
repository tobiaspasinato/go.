package main

import(
	"fmt"
)

func main(){
	var(
		edad int
	)
	fmt.Println("Cual es tu edad? ")
	fmt.Scanln(&edad)
	if(edad > 17){
		fmt.Println("tu edad es ", edad, ", sos mayor de edad")
	} else {
		if(edad > 12 && edad < 18){
			fmt.Println("tu edad es ", edad, ", sos un adolecente")
		} else {
			if(edad < 12){
				fmt.Println("tu edad es ", edad, ", sos un niño")
			}
		}
	}
}