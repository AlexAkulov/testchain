package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StateKeyPrefix is the prefix to retrieve all State
	StateKeyPrefix = "State/value/"
)

// StateKey returns the store key to retrieve a State from the index fields
func StateKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
