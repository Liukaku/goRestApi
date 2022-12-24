package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  float32 `json:"price"`
}

var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
    router.POST("/albums", postAlbums )
    router.GET("/albums/:id", getAlbumById)

    router.Run("localhost:8080")
}

func postAlbums(c *gin.Context){
    var newAlbum album

    // this declares a value and then uses it in an if statement in the same line
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context){
    id := c.Param("id")

    for _, single := range albums {
        if single.ID == id {
            c.IndentedJSON(http.StatusOK, single)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}
