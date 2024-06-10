/*ENCAPSULACION: Hacemos las propiedades privadas para evitar que 
se pueda acceder desde fuera solo mediante los metodos get y set*/

/*COMPOSICION: Permite crear estructuras que combinan los campos 
y los metodos con otras estructuras existentes*/

/*POLIMORFISMO: Es la capacidad de los objetos de diferentes clases de responder
el mismo mensaje o metodo:En Go se puede implementar el polimorfismo a traves de interfaces*/

/*INTERFACE: Mecanismo para definir un conjunto de metodos que deben ser implementados por
 otros tipos de datos. Define un contrato que especi*/
package book

import (
	"fmt"
)

/*Creamos INTERFACE y la funcion. De esta manera podemos llamar a al funcion PrintInfo() de
de una sola manera que es objeto.Print()*/
type Printable interface{
	PrintInfo()
}
func Print(p Printable){
	p.PrintInfo()
}
/*=====================*/

//creamos una estructura que es como una interface
/*la manera que nosotros tenemos de poner que un atributo es privado o publico
es ponerlo en mayuscula (publico) o en minuscula (privado)*/
type Book struct {
	title  string
	author string
	pages  int
}

/*creamos una funcion que nos va a crear objetos, 
es decir va a ser como un constructor*/
func NewBook(title string, author string, pages int) *Book{
	return &Book{
		title: 	title,
		author: author,
		pages: 	pages,
	}
}

//metodos SET
func (b *Book) SetTitle(title string){
	b.title = title
} 
func (b *Book) SetAuthor(title string){
	b.title = title
} 
func (b *Book) SetPages(title string){
	b.title = title
}

//metodos GET
func (b *Book) GetTitle() string{
	return b.title
} 
func (b *Book) GetAuthor() string{
	return b.author
} 
func (b *Book) GetPages() int{
	return b.pages
} 

//creamos una funcion para sacar por pantalla los datos del objeto Book
func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", b.title, b.author, b.pages)
}


/*======================================================================*/
/*======================================================================*/
/*Creamos una nueva estructura para un nuevo tipo de objetos, en la 
cual estamos utilizando otra estructura que teniamos antes que ees BOOK.
Al estar metiendo esta estructura BOOK dentro de otra estrucrura nosotros
podemos acceder a los metodos y propiedades de BOOK*/
type TextBook struct {
	Book
	editorial string
	level string
}

/*Creamos el constructor para este nuevo tipo de objetos*/
func NewTextBook (title, author string , pages int ,editorial string, level string) *TextBook{
	return &TextBook{
		Book: 		Book{title, author, pages},
		editorial: 	editorial,
		level: 		level,
	}
}

//creamos una funcion para sacar por pantalla los datos del objeto TextBook
func (c *TextBook) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nEditorial: %s\nLevel: %s\n", c.title, c.Book.author, c.Book.pages, c.editorial, c.level)
}