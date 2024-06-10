package main

import(
	"fmt"
)

/**Aqui estamos utilizando la propiedad 
 * interface para impirmir diferentes tipos de datos y
 * despues vamos a utilizar el tipo de dato any
 * para imprimir el resto de tipos de datos*/
func PrintList(list ...interface{}){
	for _,c := range list {
		fmt.Printf("%T - %v\n", c, c)
	}
}
func PrintListAny(list ...any){
	for _,c := range list {
		fmt.Printf("%T - %d\n", c, c)
	}
}

func main(){

	PrintList("pepe", 27, 15, true, "caratula")
	PrintListAny("pepe", 27, 15, true, "caratula")

}