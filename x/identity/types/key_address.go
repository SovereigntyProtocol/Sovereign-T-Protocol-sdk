package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AddressKeyPrefix is the prefix to retrieve all Address
	AddressKeyPrefix = "Address/value/"
)

// AddressKey returns the store key to retrieve a Address from the index fields
func AddressKey(
	owner string,
) []byte {
	var key []byte

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	return key
}
