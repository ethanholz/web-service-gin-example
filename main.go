package main

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

type album struct {
    ID      int  `json:"id"`
    Title   string  `json:"title"`
    Artist  string  `json:"artist"`
    Price   float64 `json:"price"`
}

var albums = []album {
    {ID: 1, Title: "Surf", Artist: "Surfaces", Price: 15.7},
    {ID: 2, Title: "Surf 2", Artist: "Surfaces", Price: 15.76},
}


func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func addAlbum(c *gin.Context) {
    var newAlbum album

    // Map the json received to the newAlbum object
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Always increment the ID based on length rather than
    // using specified
    if newAlbum.ID != len(albums) + 1 {
        newAlbum.ID = len(albums) + 1
    }
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
    id := c.Param("id")
    int_id, _ := strconv.Atoi(id)
    int_id -= 1
    if int_id <= len(albums) - 1 {
        c.IndentedJSON(http.StatusOK, albums[int_id])
        return
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.POST("/albums", addAlbum)
    router.GET("albums/:id", getAlbumById)

    router.Run("localhost:8080")
}
