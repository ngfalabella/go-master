package funciones

import _ "fmt"

func Recursividad(num int) int {
	if num == 0 || num == 1 {
		return num
	}
	return Recursividad(num-2) + Recursividad(num-1)
	

}