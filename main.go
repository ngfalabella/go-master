package main

import (
	"fmt"
	_ "go-master/pointer"
	"go-master/funciones"
)

var glob string 
const pi float64 = 3.14

func main() {

	
	// pointer.RetornarPointer()
	// fmt.Println("Fin del programa")
	// precio := 50
	// //valorActualizado := funciones.PrecioPorDiez(precio)

	// referencia := funciones.PrecioPuntero(&precio)

	// fmt.Println(referencia)
	

		// var indice_masa_corporal , textoRetornado = funciones.MasaCorporal(75.10 , 1.78 )

		// fmt.Println(indice_masa_corporal)
		// fmt.Println(textoRetornado)
  	fmt.Println("Inicio de Ejecucion")
		funciones.FuncionesArgumentosIlimitados(1)
		funciones.FuncionesArgumentosIlimitados(1,2)
		funciones.FuncionesArgumentosIlimitados(1,2,3)
		funciones.FuncionesArumentosIlimitadosDos("I","N","I","C","I","O" )
}