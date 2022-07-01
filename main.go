package main

import (
	"example/web-service-gin/lib"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var albums = lib.List{Head: nil}

func removeLastAlbum(c *gin.Context) {
	albums.RemoveTail()
	c.IndentedJSON(http.StatusOK, nil)
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums.ConvertToReverseArray())
}

func getHead(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums.Head.Album)
}

func addAlbum(c *gin.Context) {
	var newAlbum lib.Album

	// Map the json received to the newAlbum object
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	length := albums.Length
	// Always increment the ID based on length rather than
	// using specified
	if newAlbum.ID != length {
		newAlbum.ID = length
	}
	albums.Insert(newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	int_id, _ := strconv.Atoi(id)
	if int_id <= albums.Length {
		albumNode := albums.GetNodeById(int_id)
		if albumNode != nil {
			c.IndentedJSON(http.StatusOK, &albumNode.Album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	var album1 = lib.Album{ID: 1, Title: "Surf", Artist: "Surfaces", Price: 15.7}
	var album2 = lib.Album{ID: 2, Title: "Surf2", Artist: "Surfaces", Price: 15.7}
	albums.Insert(album1)
	albums.Insert(album2)
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", addAlbum)
	router.GET("/albums/:id", getAlbumById)
	router.GET("/albums/head", getHead)
	router.DELETE("/albums", removeLastAlbum)

	router.Run("localhost:8080")
}
