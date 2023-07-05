package id_gen

import (
	"crypto/rand"
	"math/big"
)

const ID_CHARACTERS string = "abcdefghijklmnopqrstuvwxyz"

// Generates a random string of n lowercase latin characters
func GenerateID(n int) (string, error) {
	bigIdCharacterLength := big.NewInt(int64(len(ID_CHARACTERS)))
	id := ""
	for i := 0; i < n; i++ {
		randomIndex, err := rand.Int(rand.Reader, bigIdCharacterLength)
		index := randomIndex.Int64()
		if err != nil {
			return "", err
		}
		id += string(ID_CHARACTERS[index])
	}
	return id, nil
}
