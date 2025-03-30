package base62

import (
	"testing"
)

func TestEncode(t *testing.T) {
	data := []byte("Hello, World!")
	encoded := Encode(data)

	// Check if the encoded string is not empty
	if encoded == "" {
		t.Error("Encoded string is empty")
		return
	}

	// Check if the encoded string is a valid base62 string
	for _, char := range encoded {
		if (char < '0' || char > '9') && (char < 'A' || char > 'Z') && (char < 'a' || char > 'z') {
			t.Errorf("Encoded string contains invalid character: %c", char)
			return
		}
	}
}

func TestDecode(t *testing.T) {
	data := []byte("Hello, World!")
	encoded := Encode(data)
	decoded, err := Decode(encoded)

	if err != nil {
		t.Errorf("Error decoding string: %v", err)
		return
	}

	if string(decoded) != string(data) {
		t.Errorf("Decoded string does not match original: got %s, want %s", decoded, data)
	}
}
