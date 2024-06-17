package main

import(
	"fmt"
)

type Documento struct{
	anio int
	autor string
	barcode string
	estado string
	numNormalizado string
	titulo string
}

func (d.Documento) Anio(anio int){
	d.anio = anio
}

func (d.Documento) Autor(){
    fmt.Println(d.Documento.autor)
}

func main(){
    d := Documento{
        anio: 1902,
		autor: "tobi",
		barcode: "123456789",
        estado: "activo",
        numNormalizado: "123456789",
        titulo: "El libro de tobi",
	}
}