package main

import (
	"example/apirest/repos"
	"example/apirest/repos/Sql"
	"github.com/gin-gonic/gin"
)




func main() {
	router := gin.Default()
	//is this the correct approach?
	//This for inMemory approach
	//Repo = Memory.NewInMemoryRepository()
	//this for a database approach
	repos.Repo = Sql.NewSqlRepository()

	//from here is the API
	router.GET("/", showHelloWorld)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}
