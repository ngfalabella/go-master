package types

import "fmt"

type Animal struct {
	tipoDeAnimal string
	patas        int
	vuela        bool
}

func CrearAnimal() {
	nuevoAnimal := Animal{"Cuervo", 2, true}
	fmt.Println(nuevoAnimal)
}