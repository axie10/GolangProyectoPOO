package animal

import "fmt"

type Animal interface {
	Sonido()
}

//Estructura y metodos para Perro
type Perro struct {
	Nombre string
}

func (p *Perro) Sonido() {
	fmt.Println(p.Nombre + " hace guau")
}

//Estructura y metodos para Gato
type Gato struct {
	Nombre string
}

func (p *Gato) Sonido() {
	fmt.Println(p.Nombre + " hace miau")
}

//Funcion para imprimir el sonido
func HacerSonido(animal Animal){
	animal.Sonido()

}