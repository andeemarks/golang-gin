package albums

import "errors"

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAll() []Album {
	return albums
}

func Add(newAlbum Album) Album {
	albums = append(albums, newAlbum)

	return newAlbum
}

func Get(albumId string) (foundAlbum Album, err error) {
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, album := range albums {
		if album.ID == albumId {
			return album, nil
		}
	}

	return Album{}, errors.New("album not found")
}
