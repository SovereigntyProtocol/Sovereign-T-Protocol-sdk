package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UniquekeyKeyPrefix is the prefix to retrieve all Uniquekey
	UniquekeyKeyPrefix = "Uniquekey/value/"
)

// UniquekeyKey returns the store key to retrieve a Uniquekey from the index fields
func UniquekeyKey(
	keyid string,
) []byte {
	var key []byte

	keyBytes := []byte(keyid)
	key = append(key, keyBytes...)
	key = append(key, []byte("/")...)

	return key
}
