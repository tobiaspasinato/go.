package main

import(
	"fmt"
)

func main(){
	var edad int
	var estado_civil string
	fmt.Println("Tu edad?")
	fmt.Scanln(&edad)
	fmt.Println("Tu estado civil?")
	fmt.Scanln(&estado_civil)
	if edad < 18 && estado_civil != "soltero"{
		fmt.Println("Sos muy chico para NO estar soltero")
	} else {
		fmt.Println("Tu edad es ", edad, ", y tu estado es ", estado_civil)
	}
}