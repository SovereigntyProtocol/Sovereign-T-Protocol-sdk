package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// IdKeyPrefix is the prefix to retrieve all Id
	IdKeyPrefix = "Id/value/"
)

// IdKey returns the store key to retrieve a Id from the index fields
func IdKey(
	did string,
) []byte {
	var key []byte

	didBytes := []byte(did)
	key = append(key, didBytes...)
	key = append(key, []byte("/")...)

	return key
}
