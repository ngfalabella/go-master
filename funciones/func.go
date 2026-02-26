package funciones

import (
	"fmt"
)

func PrecioPorDiez( precio int ) int {
	fmt.Println(precio , "------" , &precio)
	return precio * 10
}

func PrecioPuntero ( puntero *int )int {
	fmt.Println(puntero)
	return *puntero * 10
}