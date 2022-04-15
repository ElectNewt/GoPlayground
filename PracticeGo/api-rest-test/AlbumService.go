package main

import (
	"example/apirest/Dtos"
	"example/apirest/Services"
	"example/apirest/repos"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//var Repo repos.IRepository

//this should be split in proper services, but the logic is very simple

func showHelloWorld(c *gin.Context) {
	c.String(http.StatusOK, "%s", "hello world")
}
func getAlbums(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, repos.Repo.GetAllAlbums())
}

func postAlbums(c *gin.Context) {
	var newAlbum Dtos.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	newAlbum, error := Services.CreateAlbum(newAlbum)
	if error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": error.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	album, err := repos.Repo.GetAlbumById(idInt)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, album)

}
