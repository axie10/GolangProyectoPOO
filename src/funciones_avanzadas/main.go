package main

import "fmt"

/**FUNCIONES VARIÁDICA => Son funciones a las cuales podemos
 * meterle mas de 2 datos --"nums ...int"--
 * Tambien hay una opcion que acepta todo tipo de datos no
 * solo un tipo --"datos ...interface{}"--*/
func suma(name string, nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	fmt.Printf("Hola soy %s, y la suma es %d\n", name, total)
	return total
}

/*de estra manera recibimos datos de diferentes tipos*/
func imprimirDatos(datos ...interface{}) {
	for _, c := range datos {
		fmt.Printf("%T - %v\n", c, c)
	}
}

/**
 * FUNCIONES RECURSIVA => Funciones que se llaman a si mismas
 */
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

/**
 * FUNCIONES ANÓNIMAS => Funcion que no tiene nombre
 * Aqui tenemos una funcion que le pasamos por parametro
 * otra funcion, la funcion que le pasamos por paramtro es
 * la funcion anonima que tenemos creada dentro del main.
 */
/**Esta funcion recibe otra funcion conmo parametro */
func saludar(name string, f func(string)){
	f(name)
}
func duplicar(n int) int {
	return n * 2
}
func triplicar(n int) int {
	return n * 3
}
/**Esta funcion recibe otra funcion como parametro */
func aplicar(f func(int) int, n int) int{
	return f(n)
}

/**
 * FUNCIONES DE ORDEN SUPERIOR => Es aquella funcion que recibe por
 * parametro una funcion y ademas devuelve una funcion como resultado.
 * Permite una programacion más modular y flexible.
 */
func double(f func(int) int, x int) int{
	return f(x * 2)
}
func addOne(x int) int{
	return x + 1
}


/**
 * FUNCIONES CLOSURES => Funcion anonima que se define dentro de otra funcion
 * y puede acceder, modificar variables definidas en el ambito de la funcion 
 * externa
 */
func incrementer() func() int {
	i := 0
	return func() int{
		i++
		return i
	}
}




func main() {
	// fmt.Println(suma("suma",12,78,98,77,5,67,5,45,76))
	// imprimirDatos(23, "loca", 34, "suki", 21, true)
	// fmt.Println(factorial(5))

	/**
	 * FUNCIONES ANÓNIMAS => Funcion que no tiene nombre
	 */
	// saludo := func (name string) {
	// 	fmt.Printf("Hola %s\n", name)
	// }
	/*Le pasamos el nombre y la funcion por parametro*/
	// saludar("Pepe", saludo)
	// saludo("pepe")
	/*Aqui tambien estamos pasando funciones como parametro a otras funciones*/
	r1 := aplicar(duplicar, 3)
	r2 := aplicar(triplicar, 3)
	fmt.Println(r1, r2)
	//Ejecucion de funcion de ORDEN SUPERIOR
	r3 := double(addOne, 2)
	fmt.Println("RESULTADO: ",r3)
	/*Ejecucion de FUNCION CLOSURES, podemos ver que si
	seguimos ejecutando la funcion recuerda el valor de i
	de la funcion que se acaba de ejecutar*/
	nextInt := incrementer()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())


}
