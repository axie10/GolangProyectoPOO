package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
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
		fmt.Println(c)
	}
}

/**
 * RESTRICCIONES
 */
//esto es una restriccion creada a mano por el programador
type Numbers  interface {
	~int | ~float64 | ~float32 | ~uint
}
/** el paquete 'constrains' es el que estamos importando 
 * que tiene golang para poner directamente restricciones*/
func Sum[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _,num := range nums {
		total += num
	}
	return total
}
/**Creamos una funcion con la palabra reservada "COMPARABLE",
 * QUE ES UNA RESTRICCION, y la utilizamos para que nos permita 
 * comparar todo tipo datos
*/
func Includes[T comparable](list []T, value T) bool{
	for _,c := range list {
		if c == value {
			return true
		}
	}
	return false
}
func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T{

	result := make([]T, 0 , len(list))

	for _,item := range list{

		if callback(item) {
			result = append(result, item)
		}
	}

	return result
}

/**ESTRUCTURAS GENERICAS:*/
type Product[P uint | string] struct {
	id P
	descripcion string
	precios float32
}

func main(){

	// PrintList("pepe", 27, 15, true, "caratula")
	// PrintListAny("pepe", 27, 15, true, "caratula")

	// numeros := Sum(1,2,3,4)
	// fmt.Println(numeros)

	// strings := []string{"a", "b", "c", "d", "e"}
	// numeros1 := []int{1,2,3,4,5,6}
	// r1 := Includes(strings, "p")
	// r2 := Includes(numeros1, 4)

	// fmt.Println(r1, r2)
	// fmt.Println(Filter(numeros1, func(value int) bool {return value > 3}))
	// fmt.Println(Filter(strings, func(value string) bool {return value > "b"}))

	producto1 := Product[uint] {1, "zapatillas", 50}
	producto2 := Product[string] {"2", "zapatillas", 50}

	fmt.Println(producto1, producto2)

}