package tiny_thumb

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"testing"
)

var b64Head = mustBase64Decode(`/9j/2wCEABsSFBcUERsXFhceHBsgKEIrKCUlKFE6PTBCYFVlZF9VXVtqeJmBanGQc1tdhbWGkJ6jq62rZ4C8ybqmx5moq6QBHB4eKCMoTisrTqRuXW6kpKSkpKSkpKSkpKSkpKSkpKSkpKSkpKSkpKSkpKSkpKSkpKSkpKSkpKSkpKSkpKSkpP/AABEIABQAFAMBIgACEQEDEQH/xAGiAAABBQEBAQEBAQAAAAAAAAAAAQIDBAUGBwgJCgsQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+gEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoLEQACAQIEBAMEBwUEBAABAncAAQIDEQQFITEGEkFRB2FxEyIygQgUQpGhscEJIzNS8BVictEKFiQ04SXxFxgZGiYnKCkqNTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqCg4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2dri4+Tl5ufo6ery8/T19vf4+fr/2gAMAwEAAhEDEQA/AA==`)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// WARNING: If this test starts failing it may be because golang's jpeg library
// changed how it encodes images. (https://golang.org/src/image/jpeg/writer.go)
// if so it may require vendoring / forking this library at a previous version
// to avoid breaking images that have already been produced by this program.
func TestGolangJpegEncodingHasNotChanged(t *testing.T) {
	b := testImage(400, 400)
	b, err := minimize(b, 20, 30)
	check(err)
	head, _, err := split(b)
	check(err)
	if !bytes.Equal(head, b64Head) {
		t.Fatalf("unexpected head; got %v expected %v", head, b64Head)
	}
	d, off, err := extractDimensions(head)
	check(err)

	exp := []byte{0, 20, 0, 20}
	if !bytes.Equal(d, exp) {
		t.Fatalf("unexpected dimensions; got %v expected %v", d, exp)
	}
	expoff := 141
	if off != expoff {
		t.Fatalf("unexpected dimension offset; got %d expected %d", off, expoff)
	}
}

func TestDimensions(t *testing.T) {
	b := testImage(100, 200)
	r, err := TinyThumb(b, 1, 30, true)
	check(err)

	if r.Debug.Height != 15 || r.Debug.Width != 30 || r.Debug.Parameters.DimensionOffset != 141 {
		t.Fatalf("unexpected dimensions: %+v", r.Debug)
	}
}

func TestParameterHeaders(t *testing.T) {
	for ty := range types {
		b := testImage(100, 200)
		_, err := TinyThumb(b, ty, 30, true)
		if err != nil {
			t.Fatalf("for type %d: %v", ty, err)
		}
	}
}

func testImage(height, width int) []byte {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	cyan := color.RGBA{100, 200, 200, 0xff}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2:
				img.Set(x, y, cyan)
			case x >= width/2 && y >= height/2:
				img.Set(x, y, color.White)
			default:
				// Zero value.
			}
		}
	}

	//f, _ := os.Create("image.png")
	b := bytes.NewBuffer(nil)
	err := png.Encode(b, img)
	check(err)
	return b.Bytes()
}
