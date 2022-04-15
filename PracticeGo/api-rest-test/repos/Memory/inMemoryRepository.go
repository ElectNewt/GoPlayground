package Memory

import (
	"errors"
	"example/apirest/Dtos"
	"example/apirest/repos"
)

type InMemoryRepository struct {
	albums []Dtos.Album
}

func NewInMemoryRepository() repos.IRepository {
	repo := new(InMemoryRepository)
	repo.albums = []Dtos.Album{
		{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
	return repo
}

func (repo InMemoryRepository) GetAllAlbums() []Dtos.Album {
	return repo.albums
}

func (repo InMemoryRepository) GetAlbumById(id int) (Dtos.Album, error) {
	//this if is just for testing the error functionality
	if cap(repo.albums) >= id {
		return repo.albums[id], nil
	}
	return Dtos.Album{}, errors.New("album not found")
}
func (repo InMemoryRepository) AddAlbum(album Dtos.Album)  Dtos.Album {
	album.ID = int64(len(repo.albums)) + 1
	repo.albums = append(repo.albums, album)
	return album
}
