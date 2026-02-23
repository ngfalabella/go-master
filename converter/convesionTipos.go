package converter

import (
	"fmt"
	"strconv"
)

func ConvertirTipos() {
 todoString := "true"
 fmt.Println(todoString)
 var todoBool bool 
 todoBool , _ = strconv.ParseBool(todoString)
 fmt.Println("El resultado fue : " , todoBool)

 todoString2  := strconv.FormatBool(todoBool)

 fmt.Println(todoString2)
	

}