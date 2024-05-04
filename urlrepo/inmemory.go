package urlrepo

import (
	"errors"

	"cafebean.xyz/cafeshort/v2/id_gen"
)

const ID_LENGTH int = 6

var ErrUrlNotExists error = errors.New("url does not exist")

// Stores URL and short URL pairs in memory
type InMemoryURLRepository struct {
	storageMap map[string]string
}

// returns the longUrl of the shortId if it exists, otherwise an error
func (imus *InMemoryURLRepository) GetUrlById(shortId string) (originalUrl string, err error) {
	originalUrl, ok := imus.storageMap[shortId]
	if !ok {
		return originalUrl, ErrUrlNotExists
	}
	return originalUrl, nil
}

// Stores the long URL and returns the shortId, returns an error if this didn't work
func (imus *InMemoryURLRepository) AddUrl(longUrl string) (id string, err error) {
	//TODO: implement URL mapping
	// take url
	// generate ID
	id, err = id_gen.GenerateID(ID_LENGTH)
	if err != nil {
		return
	}
	// map the ID to the URL
	imus.storageMap[id] = longUrl
	return
}

func NewInMemoryURLRepository() *InMemoryURLRepository {
	return &InMemoryURLRepository{
		storageMap: make(map[string]string),
	}
}
