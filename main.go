/* Create the data */
package main

// NOTE: Use "()", not "{}" for import
import (
	"net/http"

	"github.com/gin-gonic/gin" // "go get ."
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums) // associate the GET method and the path with the handler
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080") // attacuh the router to the http server
}

/* a handler to return all items */
// getAlbums() responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums) // seriarize the struct into JSON and add it to thre response. StatusOK means 200 OK
}

/* a handler to add a new item */
// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJson to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to sth slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum) // status code 201
}
