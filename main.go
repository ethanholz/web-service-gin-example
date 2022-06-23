package main

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

var albums = List{nil}

type album struct {
    ID      int  `json:"id"`
    Title   string  `json:"title"`
    Artist  string  `json:"artist"`
    Price   float64 `json:"price"`
}

type Node struct {
    album album
    next *Node
}

type List struct {
    head *Node
}

func (L *List) insert(album album){
    list := &Node {
        album: album,
        next: L.head,
    }
    L.head = list
}

func (L *List) getLength() int {
    length := 0
    start := L.head
    length++
    for start != nil {
        start = start.next
        length++
    }
    return length
}

func (L *List) getNodeById(id int) *Node {
    start := L.head
    for start != nil {
        if start.album.ID == id {
            return start
        }
        start = start.next
    }
    return nil
}

func (L *List) convertToList() []album {
    var albumArray []album
    start := L.head
    for start != nil {
        albumArray = append(albumArray, start.album)
        start = start.next
    }
    return albumArray
}



func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums.convertToList())
}

func addAlbum(c *gin.Context) {
    var newAlbum album

    // Map the json received to the newAlbum object
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }
    length := albums.getLength()
    // Always increment the ID based on length rather than
    // using specified
    if newAlbum.ID != length {
        newAlbum.ID = length
    }
    albums.insert(newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
    id := c.Param("id")
    int_id, _ := strconv.Atoi(id)
    if int_id <= albums.getLength() {
        albumNode := albums.getNodeById(int_id)
        if albumNode != nil {
            c.IndentedJSON(http.StatusOK, &albumNode.album)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
    var album1 = album{ID: 1, Title: "Surf", Artist: "Surfaces", Price: 15.7}
    var album2 = album{ID: 2, Title: "Surf2", Artist: "Surfaces", Price: 15.7}
    albums.insert(album1)
    albums.insert(album2)
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.POST("/albums", addAlbum)
    router.GET("albums/:id", getAlbumById)

    router.Run("localhost:8080")
}
