package pointer

import "fmt"

func RetornarPointer() {
	var nombre_uno string = "Nicolas"
	var nombre_dos string = "Nicolas777"
	var nombre_tres string = nombre_uno
	nombre_cuatro  := &nombre_uno

	fmt.Println("1)" , nombre_uno , &nombre_uno )
	fmt.Println("2)" , nombre_dos , &nombre_dos )
	fmt.Println("3)" , nombre_tres , &nombre_tres )
	fmt.Println("4)" , *nombre_cuatro , &nombre_cuatro )
}