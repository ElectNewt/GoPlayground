package Tests

import (
	"example/apirest/Dtos"
	"example/apirest/Services"
	"example/apirest/repos"
	"example/apirest/repos/Memory"
	"testing"
)

func Test_NewAlbumGetsCreatedSuccessfully(t *testing.T) {
	repos.Repo = Memory.NewInMemoryRepository() //TODO: Mock
	album := Dtos.Album{
		ID:     0,
		Title:  "title here",
		Artist: "artist",
		Price:  10,
	}
	newAlbum, err := Services.CreateAlbum(album)

	if err != nil {
		t.Fatalf("test not successfull, id > 0 should not be able to work")
	}

	if newAlbum.ID != 4 {
		t.Fatalf("%d != %d", newAlbum.ID, 4)
	}

}

func Test_WhenTitleAlreadyExistThenError(t *testing.T) {
	repos.Repo = Memory.NewInMemoryRepository() //TODO: Mock
	album := Dtos.Album{
		ID:     0,
		Title:  "Blue Train",
		Artist: "artist",
		Price:  10,
	}
	_, err := Services.CreateAlbum(album)

	if err == nil {
		t.Fatalf("Title already exist, so this should fail")
	}
}

func Test_WhenIdIsNot0forCreateThenError(t *testing.T) {
	album := Dtos.Album{
		ID:     5,
		Title:  "title here",
		Artist: "artist",
		Price:  10,
	}

	_, err := Services.CreateAlbum(album)

	if err == nil {
		t.Fatalf("test not successfull, id > 0 should not be able to work")
	}
}
