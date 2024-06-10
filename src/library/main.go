package main

import (
	// "fmt"
	// "library/book"
	"library/animal"

)

func main() {

/*CODIGO QUE HE UTILIZADO PARA PRACTICAR LA LIBRERIA BOOK */
	//tambien se pueden crear asi los objetos
	/*Esto solo nos deja crear los objetos cuando estan las propiedades en
	en publicas, si estan en privadas no*/
	// var myBook = book.Book{
	// 	Title: "Mody Dick",
	// 	Author: "Herman Melville",
	// 	Pages: 300,
	// }

	/*asi creamos los objetos con la funcion constructor que hemos creado en el
	paquete book. Cuando tenemos en privado las propiedades del objeto es cuando
	se vuelve util el constructor*/
	// var myBook = book.NewBook("raul", "perez", 300)
	// var mytextbook = book.NewTextBook("pepe", "luis", 300, "santillana", "25")

	// myBook.PrintInfo()
	// mytextbook.PrintInfo()

	// book.Print(mytextbook)
	// book.Print(myBook)

/*CODIGO QUE HE UTILIZADO PARA PRACTICAR LA LIBRERIA ANIMAL*/

	animales := []animal.Animal {

		&animal.Perro{Nombre: "perro1"},
		&animal.Perro{Nombre: "perro2"},
		&animal.Gato{Nombre: "gato1"},
		&animal.Gato{Nombre: "gato2"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}



}