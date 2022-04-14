package main

import "errors"

type IRepository interface {
	GetAllAlbums() []album
	GetAlbumById(id int) (album, error)
	AddAlbum(album album)
}

type InMemoryRepository struct {
	albums []album
}


func NewInMemoryRepository() *InMemoryRepository {
	repo := new(InMemoryRepository)
	repo.albums = []album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
	return repo
}

func (repo InMemoryRepository) GetAllAlbums() []album {
	return repo.albums
}


func (repo InMemoryRepository) GetAlbumById(id int) (album, error) {
	//this if is just for testing the error functionality
	if cap(repo.albums) >= id {
		return repo.albums[id], nil
	}
	return album{}, errors.New("album not found")
}
func (repo InMemoryRepository) AddAlbum(album album) {
	repo.albums = append(repo.albums, album)
}
