package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var repo IRepository

func showHelloWorld(c *gin.Context) {
	c.String(http.StatusOK, "%s", "hello world")
}
func getAlbums(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, repo.GetAllAlbums())
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	repo.AddAlbum(newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	album, err := repo.GetAlbumById(idInt)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, album)

}
