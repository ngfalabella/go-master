package funciones

import (
	"fmt"
)

func EjercicioUno(x int , y int ) int {
	return x + y 
}

func EjercicioDos ( x int , y int ) ( ejemplo  int ) {
	ejemplo = x + y
	return
}

func MasaCorporal ( peso float64 , altura float64 ) (IMC float64  ,  literal string ){
	IMC = peso / (altura * altura )
	literal = "Tu indice de masa corporal es " + fmt.Sprintf("%f" , IMC)
	return
}

func FuncionesArgumentosIlimitados ( nums ... int ) {
	fmt.Println(nums , " ")
}

func FuncionesArumentosIlimitadosDos ( cadenas ... string ) {
	fmt.Println(cadenas , " ")
}