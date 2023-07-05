package url_storer

import (
	"fmt"

	"cafebean.xyz/cafeshort/v2/id_gen"
)

const ID_LENGTH int = 6

// Stores URL and short URL pairs in memory
type InMemoryURLStorer struct {
	storageMap map[string]string
}

// returns the longUrl of the shortId if it exists, otherwise an error
func (imus *InMemoryURLStorer) Query(shortId string) (originalUrl string, err error) {
	originalUrl, ok := imus.storageMap[shortId]
	if !ok {
		return originalUrl, fmt.Errorf("url does not exist")
	}
	return originalUrl, nil
}

// Stores the long URL and returns the shortId, returns an error if this didn't work
func (imus *InMemoryURLStorer) Store(longUrl string) (shortId string, err error) {
	//TODO: implement URL mapping
	// take url
	// generate ID
	shortId, err = id_gen.GenerateID(ID_LENGTH)
	if err != nil {
		return
	}
	// map the ID to the URL
	imus.storageMap[shortId] = longUrl
	// return the ID
	return shortId, nil
}

func NewInMemoryURLStorer() *InMemoryURLStorer {
	return &InMemoryURLStorer{
		storageMap: make(map[string]string),
	}
}

// Defines a type that stores URLs and their shortened version
type URLStorer interface {
	Query(shortUrl string) (string, error)
	Store(originalUrl string) (string, error)
}
