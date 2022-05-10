# Tiny Thumb 🤏
A novel, efficient, and practical method of implementing lossy compression that produces visually appealing image previews.

![](./img/tiny-thumb-ios.gif)

This technique is useful in client/server models where bandwidth is a constraint and custom client side manipulation of the payload, prior to display, is possible. Slack uses this algorithm in its product to efficiently inline image previews into some API responses. These previews are displayed to end users while higher quality images are being fetched over the network.

## Example Usage
In the example below we will downsample a 32067 byte image to 210 bytes and then upsample and post-process.

```
% identify img/cy.jpg
img/cy.jpg JPEG 256x341 256x341+0+0 8-bit sRGB 32067B 0.000u 0:00.000

% go run main.go -d 64 -o img/tiny-cy.jpg img/cy.jpg  
{
  "Debug": {
    "Parameters": {
      "Quality": 7,
      "Head": "/9j/2wCEAHJPVmRWR3JkXWSBeXKIq/+6q52dq//6/8////////////////////////////////////////////////////8BeYGBq5ar/7q6///////////////////////////////////////////////////////////////////////////AABEIAAAAAAMBIgACEQEDEQH/xAGiAAABBQEBAQEBAQAAAAAAAAAAAQIDBAUGBwgJCgsQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+gEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoLEQACAQIEBAMEBwUEBAABAncAAQIDEQQFITEGEkFRB2FxEyIygQgUQpGhscEJIzNS8BVictEKFiQ04SXxFxgZGiYnKCkqNTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqCg4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2dri4+Tl5ufo6ery8/T19vf4+fr/2gAMAwEAAhEDEQA/AA==",
      "DimensionOffset": 141
    },
    "Final": "/9j/2wCEAHJPVmRWR3JkXWSBeXKIq/+6q52dq//6/8////////////////////////////////////////////////////8BeYGBq5ar/7q6///////////////////////////////////////////////////////////////////////////AABEIAEAAMAMBIgACEQEDEQH/xAGiAAABBQEBAQEBAQAAAAAAAAAAAQIDBAUGBwgJCgsQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+gEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoLEQACAQIEBAMEBwUEBAABAncAAQIDEQQFITEGEkFRB2FxEyIygQgUQpGhscEJIzNS8BVictEKFiQ04SXxFxgZGiYnKCkqNTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqCg4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2dri4+Tl5ufo6ery8/T19vf4+fr/2gAMAwEAAhEDEQA/AIgc9SKXPP1pMc5FL3z1pAJng5/KkHtTsZOaZQA5Tx9KKUDj8M00/WgBw6nFSEY9MUYyBgYOKXJBPHTikAzhsk0pVcZUe9LwDg4GRTeWyFHFACDJzTCacchcg9aa3TP86YEpwDjpSbjxwTSbumec9qk2uF6E0gEyM4IHNN4VsflT1UhAT97pQBg5JH4U0gEC8EbTiomGKskkJ1JzUJfPbNAycsCeAPrSK3NK23byBiq+7aeKALBwSKTjd6U0MGGaC3IB6/zoASV+wqOkIwxzRQNH/9k=",
    "Height": 64,
    "Width": 48,
    "PayloadLen": 216,
    "MaxDimension": 64
  },
  "Payload": "AQBAADCIHPUilzz9aTHORS989aQCZ4OfypB7U7GTmmUAOU8fSilA4/DNNP1oAcOpxUhGPTFGMgYGDilyQTx04pAM4bJNKVXGVHvS8A4OBkU3lshRxQAgyc0wmnHIXIPWmt0z/OmBKcA46Um48cE0m7pnnPapNrhehNIBMjOCBzTeFbH5U9VIQE/e6UAYOSR+FNIBAvBG04qJhirJJCdSc1CXz2zQMnLAngD60itzStt28gYqvu2nigCwcEik43elNDBhmgtyAev86AElfsKjpCMMc0UDR//Z"
}

% convert img/tiny-cy.jpg -resize 128 -quality 100 -blur 0x4 img/tiny-cy-blur.jpg
```

| Original | Tiny | Upsampled and Blurred |
| --- | --- | --- |
| ![](./img/cy.jpg) | ![](./img/tiny-cy.jpg) | ![](./img/tiny-cy-blur.jpg) |

## High-Level Algorithim

Create:
- Convert input to JPEG.
- Downsample using predetermined JPEG quality parameters.
- Strip the JPEG header up to and including part of the 'start of scan' marker.

Reconstitute:
- Concatenate a known header with the tiny thumb and make some small adjustments to get a valid JPEG.
- Optionally upsample and post process. :sparkles:

## Options and Configuration
This program takes two options that can be changed; a type and a maximum dimension. Each type corresponds to a specific jpeg header and dimension offset. This program includes a mapping of types currently used by Slack, the details of which can be found in the program output or the source code. A server and client should preshare the mapping from all known types to corresponding headers and dimension offsets.

## Program Output and Detailed Reconstitution Algorithim
The output of this program is a json object containing:
| Key | Value |
| --- | --- |
| `Payload` | base64 encoded tiny thumb |
| `Debug.Head` | base64 encoded JPEG header |
| `Debug.DimensionOffset` | an offset at which the header must be modified (see below) |

Upon receiving a payload, a client can reconstitute a valid JPEG using the following process:
- Split payload into three parts; the first byte is the `$type`, the next four bytes are the `$dimensions`, and the remaining bytes are the `$tail`.
- Get the corresponding preshared `$header` and `$dimension_offset` for `$type`. If you do not have values for this `$type`, fail.
- Set the four bytes of `$header` beginning at offset `$dimension_offset` to the value of `$dimensions`.
- Concatenate `$header` and `$tail` to get a valid JPEG.

## Debugging Tips
- Use `cmp` to determine if the first n bytes of two files are identical.
- Tiny thumb's `-o` flag can be used to output the entire image for debugging.

## References
![jpeg-reference](./img/jpeg-reference.jpg)

The full JPEG specification: https://www.w3.org/Graphics/JPEG/jfif3.pdf

Discussion and prior-art related to this technique:
- https://stackoverflow.com/questioins/56236805/create-jpeg-thumb-image-with-general-fixed-header
- https://engineering.fb.com/android/the-technology-behind-preview-photos
