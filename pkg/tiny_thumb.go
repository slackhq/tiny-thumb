package tiny_thumb

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"image"
	"image/jpeg"

	"github.com/disintegration/imaging"
)

var (
	types = map[byte]TypeParameters{
		1: {
			7, mustBase64Decode(`/9j/2wCEAHJPVmRWR3JkXWSBeXKIq/+6q52dq//6/8////////////////////////////////////////////////////8BeYGBq5ar/7q6///////////////////////////////////////////////////////////////////////////AABEIAAAAAAMBIgACEQEDEQH/xAGiAAABBQEBAQEBAQAAAAAAAAAAAQIDBAUGBwgJCgsQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+gEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoLEQACAQIEBAMEBwUEBAABAncAAQIDEQQFITEGEkFRB2FxEyIygQgUQpGhscEJIzNS8BVictEKFiQ04SXxFxgZGiYnKCkqNTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqCg4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2dri4+Tl5ufo6ery8/T19vf4+fr/2gAMAwEAAhEDEQA/AA==`), 141,
		},
		2: {
			10, mustBase64Decode(`/9j/2wCEAFA3PEY8MlBGQUZaVVBfeMiCeG5uePWvuZHI//////////////////////////////////////////////////8BVVpaeGl464KC6//////////////////////////////////////////////////////////////////////////AABEIAAAAAAMBIgACEQEDEQH/xAGiAAABBQEBAQEBAQAAAAAAAAAAAQIDBAUGBwgJCgsQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+gEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoLEQACAQIEBAMEBwUEBAABAncAAQIDEQQFITEGEkFRB2FxEyIygQgUQpGhscEJIzNS8BVictEKFiQ04SXxFxgZGiYnKCkqNTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqCg4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2dri4+Tl5ufo6ery8/T19vf4+fr/2gAMAwEAAhEDEQA/AA==`), 141,
		},
		3: {
			15, mustBase64Decode(`/9j/2wCEADUlKC8oITUvKy88OTU/UIVXUElJUKN1e2GFwarLyL6qurfV8P//1eL/5re6////////////zv////////////8BOTw8UEZQnVdXnf/cutz////////////////////////////////////////////////////////////////////AABEIAAAAAAMBIgACEQEDEQH/xAGiAAABBQEBAQEBAQAAAAAAAAAAAQIDBAUGBwgJCgsQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+gEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoLEQACAQIEBAMEBwUEBAABAncAAQIDEQQFITEGEkFRB2FxEyIygQgUQpGhscEJIzNS8BVictEKFiQ04SXxFxgZGiYnKCkqNTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqCg4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2dri4+Tl5ufo6ery8/T19vf4+fr/2gAMAwEAAhEDEQA/AA==`), 141,
		},
		4: {
			20, mustBase64Decode(`/9j/2wCEACgcHiMeGSgjISMtKygwPGRBPDc3PHtYXUlkkYCZlo+AjIqgtObDoKrarYqMyP/L2u71////m8H////6/+b9//gBKy0tPDU8dkFBdviljKX4+Pj4+Pj4+Pj4+Pj4+Pj4+Pj4+Pj4+Pj4+Pj4+Pj4+Pj4+Pj4+Pj4+Pj4+Pj4+Pj4+P/AABEIAAAAAAMBIgACEQEDEQH/xAGiAAABBQEBAQEBAQAAAAAAAAAAAQIDBAUGBwgJCgsQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+gEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoLEQACAQIEBAMEBwUEBAABAncAAQIDEQQFITEGEkFRB2FxEyIygQgUQpGhscEJIzNS8BVictEKFiQ04SXxFxgZGiYnKCkqNTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqCg4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2dri4+Tl5ufo6ery8/T19vf4+fr/2gAMAwEAAhEDEQA/AA==`), 141,
		},
		5: {
			25, mustBase64Decode(`/9j/2wCEACAWGBwYFCAcGhwkIiAmMFA0MCwsMGJGSjpQdGZ6eHJmcG6AkLicgIiuim5woNqirr7EztDOfJri8uDI8LjKzsYBIiQkMCowXjQ0XsaEcITGxsbGxsbGxsbGxsbGxsbGxsbGxsbGxsbGxsbGxsbGxsbGxsbGxsbGxsbGxsbGxsbGxv/AABEIAAAAAAMBIgACEQEDEQH/xAGiAAABBQEBAQEBAQAAAAAAAAAAAQIDBAUGBwgJCgsQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+gEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoLEQACAQIEBAMEBwUEBAABAncAAQIDEQQFITEGEkFRB2FxEyIygQgUQpGhscEJIzNS8BVictEKFiQ04SXxFxgZGiYnKCkqNTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqCg4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2dri4+Tl5ufo6ery8/T19vf4+fr/2gAMAwEAAhEDEQA/AA==`), 141,
		},
		6: {
			30, mustBase64Decode(`/9j/2wCEABsSFBcUERsXFhceHBsgKEIrKCUlKFE6PTBCYFVlZF9VXVtqeJmBanGQc1tdhbWGkJ6jq62rZ4C8ybqmx5moq6QBHB4eKCMoTisrTqRuXW6kpKSkpKSkpKSkpKSkpKSkpKSkpKSkpKSkpKSkpKSkpKSkpKSkpKSkpKSkpKSkpKSkpP/AABEIAAAAAAMBIgACEQEDEQH/xAGiAAABBQEBAQEBAQAAAAAAAAAAAQIDBAUGBwgJCgsQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+gEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoLEQACAQIEBAMEBwUEBAABAncAAQIDEQQFITEGEkFRB2FxEyIygQgUQpGhscEJIzNS8BVictEKFiQ04SXxFxgZGiYnKCkqNTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqCg4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2dri4+Tl5ufo6ery8/T19vf4+fr/2gAMAwEAAhEDEQA/AA==`), 141,
		},
	}
)

type TypeParameters struct {
	Quality         int
	Head            []byte
	DimensionOffset int
}

type (
	Debug struct {
		Parameters   TypeParameters
		Final        []byte
		Height       uint16
		Width        uint16
		PayloadLen   int
		MaxDimension int
	}
	Result struct {
		Debug   Debug
		Payload []byte
	}
)

func TinyThumb(b []byte, typeID byte, maxDimension int, checkParametersMatch bool) (*Result, error) {
	p, ok := types[typeID]
	if !ok {
		return nil, fmt.Errorf("unknown type id: %d", typeID)
	}
	buf := bytes.NewBuffer([]byte{typeID})

	min, err := minimize(b, maxDimension, p.Quality)
	if err != nil {
		return nil, fmt.Errorf("minimize: %v", err)
	}
	mincopy := make([]byte, len(min))
	copy(mincopy, min)

	head, tail, err := split(min)
	if err != nil {
		return nil, fmt.Errorf("split: %v", err)
	}
	dimen, dOff, err := extractDimensions(head)
	if err != nil {
		return nil, fmt.Errorf("extract dimensions: %v", err)
	}
	buf.Write(dimen)
	buf.Write(tail)
	for i := 0; i < 4; i++ {
		head[dOff+i] = 0
	}
	if checkParametersMatch {
		if !bytes.Equal(head, p.Head) {
			enc := base64.StdEncoding
			return nil, fmt.Errorf("unexpected header: got '%s' expected '%s'", enc.EncodeToString(head), enc.EncodeToString(p.Head))
		}
		if dOff != p.DimensionOffset {
			return nil, fmt.Errorf("unexpected dimension offset: got %d expected %d", dOff, p.DimensionOffset)
		}
	}

	return &Result{
		Debug: Debug{
			Parameters:   p,
			Final:        mincopy,
			Height:       binary.BigEndian.Uint16(dimen),
			Width:        binary.BigEndian.Uint16(dimen[2:4]),
			PayloadLen:   buf.Len(),
			MaxDimension: maxDimension,
		},
		Payload: buf.Bytes(),
	}, nil
}

// minimize converts an image to jpeg, scales it down, and sets it's 'quality'.
func minimize(b []byte, maxDimen, quality int) ([]byte, error) {
	img, t, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("unable to decode image (%s): %v", t, err)
	}
	width := img.Bounds().Max.X - img.Bounds().Min.X
	height := img.Bounds().Max.Y - img.Bounds().Min.Y
	if width >= height && width > maxDimen {
		height = int(float64(height) / float64(width) * float64(maxDimen))
		width = maxDimen
	} else if height > maxDimen {
		width = int(float64(width) / float64(height) * float64(maxDimen))
		height = maxDimen
	}
	img = imaging.Resize(img, width, height, imaging.Lanczos)
	if _, ok := img.(*image.Gray); ok {
		// The headers for grayscale images are different. TODO convert to color.
		// https://golang.org/src/image/jpeg/writer.go#L511
		return nil, fmt.Errorf("grayscale image unsupported")
	}
	outb := bytes.NewBuffer(nil)
	err = jpeg.Encode(outb, img, &jpeg.Options{
		Quality: quality,
	})
	if err != nil {
		err = fmt.Errorf("unable to encode image: %v", err)
	}
	return outb.Bytes(), err
}

// Finds the SOF0 marker, pulls out the images dimensions and the offset at
// which the dimensions are written.
func extractDimensions(b []byte) ([]byte, int, error) {
	offset := 0
	for {
		m, skip, err := readJpegMarker(b[offset:])
		if err != nil {
			return nil, 0, err
		}
		if m == sof0Marker {
			offset += 5 // Skip past the 4 byte header and the data precision byte.
			if offset+4 >= len(b) {
				return nil, 0, fmt.Errorf("EOF before dimensions")
			}
			// Make a copy of the dimen bytes instead of returning a byte slice,
			// other parts of this program may overwrite the original dimensions.
			dimen := make([]byte, 4)
			copy(dimen, b[offset:offset+4])
			return dimen, offset, nil
		}
		offset += skip
		if offset >= len(b) {
			return nil, 0, fmt.Errorf("EOF before SOF0 marker found")
		}
	}
}

// split takes a valid JPEG and splits it just after the SOS (start of scan)
// marker. In theory, everything before this point is the same for all images
// with the same dimensions.
func split(b []byte) ([]byte, []byte, error) {
	offset := 0
	for {
		m, skip, err := readJpegMarker(b[offset:])
		if err != nil {
			return nil, nil, err
		}
		if m == sosMarker {
			// Skip the marker header
			offset += 4
			// And the YcbCr components
			// https://golang.org/src/image/jpeg/writer.go?#L503
			offset += 10 // TODO maybe be smarter than hardcoding this?
			if offset >= len(b) {
				return nil, nil, fmt.Errorf("EOF before SOS payload")
			}
			return b[:offset], b[offset:], nil
		}
		offset += skip
		if offset >= len(b) {
			return nil, nil, fmt.Errorf("EOF before SOS marker found")
		}
	}
}

const (
	soiMarker  = 0xd8 // Start of image.
	rst0Marker = 0xd0 // ReSTart (0).
	rst7Marker = 0xd7 // ReSTart (7).
	eoiMarker  = 0xd9 // End Of Image.
	sof0Marker = 0xc0 // Start Of Frame 0
	sosMarker  = 0xda // Start Of Scan.
	driMarker  = 0xdd // Define Restart Interval.
)

func readJpegMarker(b []byte) (byte, int, error) {
	if len(b) < 2 {
		return 0, 0, fmt.Errorf("buffer underflow while parsing jpeg marker")
	}
	if b[0] != 0xFF {
		return 0, 0, fmt.Errorf("bad marker prefix")
	}
	if b[1] == soiMarker || b[1] == eoiMarker || (b[1] >= rst0Marker && b[1] <= rst7Marker) {
		return b[1], 2, nil
	}

	if len(b) < 4 {
		return 0, 0, fmt.Errorf("buffer underflow while parsing jpeg marker")
	}
	if b[1] == driMarker {
		return b[1], 4, nil
	}

	// At this point we are pretty sure this a variable length header.
	n := int(b[2])<<8 + int(b[3]) - 2
	if n < 0 {
		return 0, 0, fmt.Errorf("bad length")
	}
	return b[1], n + 4, nil
}

func mustBase64Decode(s string) []byte {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}
