package encoding

import (
	"fmt"
)

var ASCII Encoder = &asciiEncoder{}

type asciiEncoder struct{}

func (e asciiEncoder) Encode(data []byte) ([]byte, error) {
	var out []byte
	for _, r := range data {
		if r > 127 {
			return nil, fmt.Errorf("invalid ASCII char: '%s'", string(r))
		}
		out = append(out, r)
	}

	return out, nil
}

func (e asciiEncoder) Decode(data []byte, length int) ([]byte, int, error) {
	// read only 'length' bytes (1 byte - 1 ASCII character)
	if len(data) < length {
		return nil, 0, fmt.Errorf("not enough data to decode. expected len %d, got %d", length, len(data))
	}
	data = data[:length]
	var out []byte
	for _, r := range data {
		if r > 127 {
			return nil, 0, fmt.Errorf("invalid ASCII char: '%s'", string(r))
		}
		out = append(out, r)
	}

	return out, length, nil
}
