package types

import (
	"fmt"
)

type Persona struct {
	name string
	age int 
	email string
	direccion Direccion
}


type Direccion struct {
	calle string 
	numero int
}

func RetornarPersona () {
	dir := Direccion{"Calle Falsa" , 123 }
	personaje_uno := Persona{"Eduardo",123,"eduardo@gmail.com", dir}
	fmt.Println(personaje_uno)
	personaje_uno.SaludarNombre()
}

func ( p *Persona) SaludarNombre() {
	fmt.Printf("Hola soy %s tengo %d y vivo en %s %d" , p.name , p.age , p.direccion.calle,  p.direccion.numero)
}