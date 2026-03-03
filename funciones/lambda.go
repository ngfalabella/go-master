package funciones

import(
	_ "fmt"
)

func LambdaFuncion(lado float64) (area func() float64) {
	area = func() float64 {
		return lado * lado
	}
	return
}