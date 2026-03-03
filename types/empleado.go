package types
import(
	"fmt"
)


type Empleado struct {
	nombre string
	email string
	direccion string
	tecnologia [3]Tecnologias
	experiencia int
}

type Tecnologias struct{
	nombre string
	experiencia string
}

func MostrarCurriculum () {
	var tecno [3]Tecnologias 

	tecno[0].nombre = "React"
	tecno[0].experiencia = "Mid"

	tecno[1].nombre = "Golang"
	tecno[1].experiencia = "Expert"

	tecno[2].nombre = "Javascript"
	tecno[2].experiencia = "Junior"


	cur := Empleado{nombre : "Nicolas " , email : "nicolas@falabella.com" , direccion : "160 de piola" , tecnologia :tecno , experiencia : 2}
	fmt.Println(cur)
}