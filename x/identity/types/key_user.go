package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UserKeyPrefix is the prefix to retrieve all User
	UserKeyPrefix = "User/value/"
)

// UserKey returns the store key to retrieve a User from the index fields
func UserKey(
	did string,
) []byte {
	var key []byte

	didBytes := []byte(did)
	key = append(key, didBytes...)
	key = append(key, []byte("/")...)

	return key
}
