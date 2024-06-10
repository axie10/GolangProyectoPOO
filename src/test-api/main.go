package main

/**GO RUNTIMES =>  Son una forma de realizar concurrencia en Go,
 * permitiendo que múltiples tareas se ejecuten de manera concurrente dentro del
 * mismo programa. Una goroutine es una función o método que se ejecuta
 * concurrentemente con otras goroutines en el mismo espacio de direcciones.
 * CANALES => Los canales son un caracteristica fundamental para la
 * comunicacion y sincronizacion de goruntimes, es decir entre sus
 * procesos ligeros, actuando como un conductos por el cual fluye la
 * informacion. Un canal se declara usando la funcion make().
 */

import (
	"fmt"
	"net/http"
	"time"
	// "time"
)

func main(){

	start := time.Now()

	apis := []string {
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://api.somewhereintheinternet.com",
		"https://graph.microsoft.com",
	}

	ch := make(chan string)

	for _,url := range apis {
		/*con el "go" estamos aplicando el "gorountime"
		lo que hace es que va haciendo las comprobaiones
		a la vez, y la que menos tiempo tarda en terminar
		es la primera*/
		go checkAPI(url, ch)
	}

	for i:=0; i<len(apis); i++{
		fmt.Print(<-ch)
	}

	alapsed := time.Since(start)
	fmt.Printf("Completado, ha tardado %s segundos", alapsed)
}

//Metodo para saber si una apie sta en funcionamiento
func checkAPI(api string, ch chan string){
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s\n", api)
		return
	}
	
	ch <- fmt.Sprintf("La api %s esta en funcionamiento\n", api)
}