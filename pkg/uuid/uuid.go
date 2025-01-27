// This package replaces the uuid package from google
package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
)

// UUID represents a 128-bit UUID value
type UUID [16]byte

var Nil UUID

// New generates a new random UUID
func New() UUID {
	var uuid UUID
	rand.Read(uuid[:])
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	return uuid
}

func NewRandom() UUID {
	return New()
}

// Parse parses a UUID string
func Parse(s string) (UUID, error) {
	var uuid UUID

	// Remove dashes from the UUID string
	s = strings.ReplaceAll(s, "-", "")

	// Ensure the sanitized string is 32 characters long
	if len(s) != 32 {
		return uuid, fmt.Errorf("invalid UUID format")
	}

	// Decode the hex string into bytes
	b, err := hex.DecodeString(s)
	if err != nil {
		return uuid, err
	}
	copy(uuid[:], b)
	return uuid, nil
}

// String returns the UUID string representation
func (uuid UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

// MarshalJSON marshals the UUID to JSON
func (uuid UUID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, uuid.String())), nil
}

// UnmarshalJSON unmarshals a JSON-encoded UUID string
func (uuid *UUID) UnmarshalJSON(data []byte) error {
	// Trim the surrounding quotes from the JSON string
	s := strings.Trim(string(data), `"`)

	// Parse the UUID string
	parsedUUID, err := Parse(s)
	if err != nil {
		return err
	}

	// Copy the parsed UUID into the current UUID
	*uuid = parsedUUID
	return nil
}
