package main

import "fmt"

// import "fmt"

/**Funicon VARIÁDICA => Son funciones a las cuales podemos
 * meterle mas de 2 datos --"nums ...int"--
 * Tambien hay una opcion que acepta todo tipo de datos no 
 * solo un tipo --"datos ...interface{}"--*/
func suma(name string, nums ...int) int {
	var total int
	for _,num:= range nums {
		total += num
	}
	fmt.Printf("Hola soy %s, y la suma es %d\n", name, total)
	return total
}
func imprimirDatos(datos ...interface{}){
	for _,c := range datos {
		fmt.Println(c)
	}
}

func main () {
	fmt.Println(suma("suma",12,78,98,77,5,67,5,45,76))
	imprimirDatos(23, "loca", 34, "suki", 21)
}