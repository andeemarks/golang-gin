package albums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultAlbumsAreAccessible(t *testing.T) {
	defaultAlbums := GetAll()

	assert.Equal(t, 3, len(defaultAlbums))
}

func TestKnownAlbumsAreRetrievableViaID(t *testing.T) {
	defaultAlbums := GetAll()

	defaultAlbumId := defaultAlbums[0].ID
	album, err := Get(defaultAlbumId)

	if err != nil {
		t.Errorf("Could not find default album with key %s", defaultAlbumId)
	}

	assert.Equal(t, defaultAlbums[0], album)
}

func TestUnknownAlbumsCannotBeRetrieved(t *testing.T) {
	_, err := Get("not-an-album-id")

	if err == nil {
		t.Errorf("Found unexpected album with key %s", "not-an-album-id")
	}
}

func TestAddingNewAlbumsIncreasedTheAlbumCount(t *testing.T) {
	startingAlbumCount := len(GetAll())

	Add(Album{ID: "NewId", Title: "New Title", Artist: "New Artist", Price: 6.66})

	endingAlbumCount := len(GetAll())
	assert.Equal(t, startingAlbumCount+1, endingAlbumCount)
}

func TestNewlyAddedAlbumsCanBeRetrieved(t *testing.T) {
	newAlbum := Album{ID: "NewId", Title: "New Title", Artist: "New Artist", Price: 6.66}
	Add(newAlbum)

	foundAlbum, _ := Get(newAlbum.ID)
	assert.Equal(t, newAlbum, foundAlbum)
}
