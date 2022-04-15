package main

import (
	"example/apirest/repos/Sql"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//is this the correct approach?
	//This for inMemory approach
	//repo = Memory.NewInMemoryRepository()
	//this for a database approach
	repo = Sql.NewSqlRepository()
	//from here works
	router.GET("/", showHelloWorld)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}
