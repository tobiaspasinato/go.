package main

import(
	"fmt"
)

type Persona struct{
	nombre string
	apellido string
	edad int
}

type Player struct{
	Persona
	gameFav string
	jugoAlLol bool
}

func (p Persona) saludar(){
    fmt.Println("Hola, soy", p.nombre, p.apellido)
}


func (p Player) jugar(){
	if p.jugoAlLol == true{
        fmt.Println(p.nombre, "jugo al lol")
    }else{
        fmt.Println(p.nombre, "no jugo al lol")
    }
}
func main(){
	jugador := Player{
		Persona: Persona{
            nombre: "Juan",
            apellido: "Perez",
            edad: 25,
        },
        gameFav: "League of Legends",
        jugoAlLol: true,
	}
	jugador.saludar()
	jugador.jugar()
}