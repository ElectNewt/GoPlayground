package Sql

import (
	"database/sql"
	"example/apirest/Dtos"
	"example/apirest/repos"
	_ "github.com/go-sql-driver/mysql"
)

type SqlRepository struct {
	Db *sql.DB
}

func NewSqlRepository() repos.IRepository {
	repo := new(SqlRepository)
	//I am not sure about this, because we should open/close the connection per request.
	//and with this solution I am opening a connection once and never closing it.
	repo.Db = GetDb()
	return repo
}

func (repo SqlRepository) GetAlbumById(id int) (Dtos.Album, error) {

	rows, error := repo.Db.Query("select * from Album where Id = ?", id)
	if error != nil {
		panic("there is an error: " + error.Error())
	}
	var album Dtos.Album
	//is it possible to map automatically?
	for rows.Next() {
		album = MapToAlbum(rows)
	}
	return album, error
}

//there should be a "createAlbumRequest" type but for simplicty i used the same.
func (repo SqlRepository) AddAlbum(album Dtos.Album) Dtos.Album {
	query := "INSERT INTO Album(Title, Artist, Price) " +
		"VALUES (?, ?, ?)"
	result, error := repo.Db.Exec(query, album.Title, album.Artist, album.Price)

	if error != nil {
		panic(error.Error())
	}
	lastId, _ := result.LastInsertId()
	album.ID = lastId

	return album
}

func (repo SqlRepository) GetAllAlbums() []Dtos.Album {

	var albums []Dtos.Album

	rows, error := repo.Db.Query("select * from Album")

	if error != nil {
		panic("there is an error: " + error.Error())
	}

	for rows.Next() {
		album := MapToAlbum(rows)
		albums = append(albums, album)
	}
	return albums
}

func MapToAlbum(row *sql.Rows) Dtos.Album {
	var album Dtos.Album
	row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)

	return album
}
