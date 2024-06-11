package main

import (
	// "fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	// "golang.org/x/text/message"
)
/**Creamos la interface de album */
type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price string `json:"price"`
}
/**Creamos la "bbdd" de albums */
var albums = []album{
	{ID: "1", Title:"Blue", Artist: "Pepe", Price: "54"},
	{ID: "2", Title:"Green", Artist: "Luis", Price: "44"},
	{ID: "3", Title:"Yellow", Artist: "Alberto", Price: "34"},
}
/**CONTROLADOR: Creamos la funcion para sacar los albumes */
func getAlbums(c *gin.Context){
	//para convertir a JSON los albumes
	c.IndentedJSON(http.StatusOK, albums)
}
/**CONTROLADOR: Creamos la funcion para guardar albumes */
func postAlbums(c *gin.Context){
	var newAlbum album
	/**BindJSON para convertir a JSON el nuevo 
	 * elemento que vamos a meter*/
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
/**CONTROLADOR: Funcion para devolver solo un album pasadnole el id */
func getAlbumsByID(c *gin.Context){
	id := c.Param("id")
	for _,a := range albums {
		if a.ID == id{
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"mensaje":"album no encontrado"})
}
/**CONTROLADOR: Borrar album */
func getAlbumsDelete(c *gin.Context){
	id := c.Param("id")
	for i, a := range albums {
		if a.ID == id{
			//Eliminamos el album
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"mensaje":"album eliminado"})
			return
			
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"mensaje":"album no encontrado"})
}
/**CONTROLADOR: Funcion para editar albumes*/
// func getEditAlbum(c *gin.Context){
// 	id := c.Param("id")
// 	title := c.Query("title")
// 	artist := c.Query("artist")
// 	price := c.Query("price")
// 	for _, a := range albums {
// 		if a.ID == id{
// 			var editAlbum album
// 			if err := c.BindJSON(&editAlbum); err != nil {
// 				return
// 			}
// 			a.Title = title
// 			a.Artist = artist
// 			a.Price = price
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"mensaje":"album no encontrado"})
// }

func main() {

	router := gin.Default()
	/**Los endpoint */
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumsByID)
	// router.PUT("/albums/:id", getEditAlbum)
	router.DELETE("/albums/:id", getAlbumsDelete)
	router.Run("localhost:8080")


}