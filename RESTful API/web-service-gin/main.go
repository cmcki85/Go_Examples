package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album structure
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Seed data for the album struct
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}

// getAlbums responds with a list of albums in JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// PostAlbums adds an album onto the request body of the JSON
func postAlbums(c *gin.Context) {
	var newAlbum album

	// call Blind JSON to get JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// If successful add to the album slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumbyID locates an album based on ID
func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	// loop over the albums and return one that matches an ID
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found!"})
}
