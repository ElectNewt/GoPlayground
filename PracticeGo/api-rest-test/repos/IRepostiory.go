package repos

import (
	"example/apirest/Dtos"
)

type IRepository interface {
	GetAllAlbums() []Dtos.Album
	GetAlbumById(id int) (Dtos.Album, error)
	AddAlbum(album Dtos.Album)  Dtos.Album
}

