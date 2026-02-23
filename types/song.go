package types

import (
	"fmt"
)

type SongType struct{
	Nombre string
	Artista string
	Orden int 
}


func ( s * SongType ) ReproducirCancion() string {
	fmt.Printf( "La cancion se llama %s , la canta %s y esta en el puesto Nº %d de la lista" , s.Nombre , s.Artista , s.Orden )
	return "Cargada con exito"
}

func ListadoCanciones () {
	//song_uno := SongType{"Jijiji" , "Patricio rey y sus redonditos de ricota" , 10}

}