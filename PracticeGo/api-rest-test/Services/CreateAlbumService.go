package Services

import (
	"errors"
	"example/apirest/Dtos"
	"example/apirest/repos"
)

//I created this to practice unit testing in go
//basically I'm going to check if the title exist before inserting,
//if it does I will return an error

func CreateAlbum(album Dtos.Album) (Dtos.Album, error) {
	if album.ID != 0 {
		return Dtos.Album{}, errors.New("Id must be 0 for creation")
	}
	allAlbums := repos.Repo.GetAllAlbums()

	for _, existingAlbum := range allAlbums {
		if album.Title == existingAlbum.Title {
			return Dtos.Album{}, errors.New("The title already exist")
		}
	}
	return repos.Repo.AddAlbum(album), nil
}
