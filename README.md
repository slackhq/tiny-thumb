# Tiny Thumb
Takes an image as input and produces a tiny payload that can be used to reconstitute a thumbnail of the input using preshared, static, image quality parameters. This method is particularly useful in client/server models where payload size is a concern and custom client side manipulation of the payload, prior to display, is possible.

At the time of writing Slack uses this algorithim in it's product to inline some image content into API responses.

## High Level Generation Algorithim Description
- Convert the image to a jpeg.
- Scale it down.
- Reduce it's 'quality' using predetermined, hardcoded, static parameters.
- Strip the jpeg header.

## Options and Configuration
There are two options that can be changed; the 'type' and the 'maximum dimension'. Each type corresponds to a specific jpeg header and 'dimension offset'.

## Output and Reconstitution Algorithim
The output of this program is a json object containing the key "Payload" whose value is a base64 encoded byte array. The base64 header can be found in the key "Debug.Head", and the dimension offset in Debug.DimensionOffset.

It is expected that a server and client will preshare the mapping from all known types to there corresponding headers and dimension offsets. Upon receiving a payload a client can reconstitute the full JPEG by using the following process:

- Split payload into three parts; the first byte is the `$type`, the next four bytes are the `$dimensions`, and the remaining bytes are the `$tail`.
- Get the corresponding preshared `$header` and `$dimension_offset` for `$type`. If you do not have values for this `$type`, fail.
- Set the four bytes of `$header` beginning at offset `$dimension_offset` to the value of `$dimensions`.
- Concatenate `$header` and `$tail` to get the final result.
- Optional: apply a gaussian blur.

## Example Usage

```
> go run main.go -d 128 -o img/tiny-cy.jpg img/cy.jpg
{
  "Debug": {
    "Parameters": {
      "Quality": 7,
      "Head": "/9j/2wCEAHJPVmRWR3JkXWSBeXKIq/+6q52dq//6/8////////////////////////////////////////////////////8BeYGBq5ar/7q6///////////////////////////////////////////////////////////////////////////AABEIAAAAAAMBIgACEQEDEQH/xAGiAAABBQEBAQEBAQAAAAAAAAAAAQIDBAUGBwgJCgsQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+gEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoLEQACAQIEBAMEBwUEBAABAncAAQIDEQQFITEGEkFRB2FxEyIygQgUQpGhscEJIzNS8BVictEKFiQ04SXxFxgZGiYnKCkqNTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqCg4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2dri4+Tl5ufo6ery8/T19vf4+fr/2gAMAwEAAhEDEQA/AA==",
      "DimensionOffset": 141
    },
    "Final": "/9j/2wCEAHJPVmRWR3JkXWSBeXKIq/+6q52dq//6/8////////////////////////////////////////////////////8BeYGBq5ar/7q6///////////////////////////////////////////////////////////////////////////AABEIAIAAYAMBIgACEQEDEQH/xAGiAAABBQEBAQEBAQAAAAAAAAAAAQIDBAUGBwgJCgsQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+gEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoLEQACAQIEBAMEBwUEBAABAncAAQIDEQQFITEGEkFRB2FxEyIygQgUQpGhscEJIzNS8BVictEKFiQ04SXxFxgZGiYnKCkqNTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqCg4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2dri4+Tl5ufo6ery8/T19vf4+fr/2gAMAwEAAhEDEQA/AIww5LE5pM9DSU4FcghTikMGwxyOlLnLYwB603J65oxnpRYLik8YBzQT0O360nRs0OQeRSAbg4z2pwJ247U2lIwO1MQq9adk7v8AGmAkHinHlckmgA3cUlAGRkUc47cUAIf1pV4PTNL2+7ijadvagBCQT3pVAJ6mjr2GPWkxk89KAFYHp6U0D2p3fGaUjnHT3oAjpfrR9KSgByjPbPvSt+VJj5c085CA85oAj+tJTt2QcimmgQ8g574pRt6L09aaxOKeFBUNx70DHYC8dR1oBDE4x0peMYHSk+6pHbtSARhnHcU0ryMd6OrgU5m4GR0NAB5OD14puOcjgU8H19KawATt9aAEODk8g+1BOQeoI4NNxkinkHqpwKBjASRx25pppSSCaSmIsKFKjjNJwgGOQfWhcgc4zSgZx7UgDjGM0uRjHamuPnUnvQRuxigAKg9O3vQwyoHWlIUKeecUiOFXn1oAToRnvTWOVFIxLMQOBT/KIXO7mmBH06ZxSlwG4Oc05MFcEA+1MfgnjAzQAZOMsOKYaeCWHzZIFNbrxTAmwQM9R/KjIycHilbOMj8aaH4z36fhUgLu+UULkkc4poU529B60m1i3vQBKcBgevNBUcknvSEAAZyfekU5yKAGYKng1IM7uHyPSkkA25B5700MTgAc0xgFwaeVDAj3o2NkYB96dtfOdtIRAQBkZpDjFSMuE78mmOMNwQRTAfvI5A49KTG4+lIMbSGGGFJlmIAHNAEoZsYb6UMTtyOCP1p6xfL8x5NNaJh0O4fypAMRt4296kKhKIowrE/lSsNz00BESD0Gach29AM04hQcdaC2OgP5UwHqzE84/Wh27UJ0yRTGoAN1MO00wsTSY9aRSRdKq3UA/WmYUH5QBSF8nApM4OO9Ahd1PB4qHOelOSgCTtxTFXB+bJJpc0oOeD1oAQqAM00Zz2NKxwelKGFMQrHAxUEjdvWnu45PpUPU5NJjQAUtFFIos7VA+6Ka6BunWnetRsxCgimSR4w3U0qvg4zTSc0zoc0AWjyM9jSc/iOnvUaPj6U/cDQMeCGHBpmc8EcjrSNkcimq5Lj17+9AhjUClkGHPpSUMaFpKKKQz//Z",
    "Height": 128,
    "Width": 96,
    "PayloadLen": 832,
    "MaxDimension": 128
  },
  "Payload": "AQCAAGCMMOSxOaTPQ0lOBXIIU4pDBsMcjpS5y2MAetNyeuaMZ6UWC4pPGAc0E9Dt+tJ0bNDkHkUgG4OM9qcCduO1NpSMDtTEKvWnZO7/ABpgJB4px5XJJoAN3FJQBkZFHOO3FACH9aVeD0zS9vu4o2nb2oAQkE96VQCepo69hj1pMZPPSgBWB6elNA9qd3xmlI5x096AI6X60fSkoAcoz2z70rflSY+XNPOQgPOaAI/rSU7dkHIppoEPIOe+KUbei9PWmsTinhQVDce9Ax2AvHUdaAQxOMdKXjGB0pPuqR27UgEYZx3FNK8jHejq4FOZuBkdDQAeTg9eKbjnI4FPB9fSmsAE7fWgBDg5PIPtQTkHqCODTcZIp5B6qcCgYwEkcduaaaUkgmkpiLChSo4zScIBjkH1oXIHOM0oGce1IA4xjNLkYx2prj51J70EbsYoACoPTt70MMqB1pSFCnnnFIjhV59aAE6EZ701jlRSMSzEDgU/yiFzu5pgR9OmcUpcBuDnNOTBXBAPtTH4J4wM0AGTjLDimGnglh82SBTW68UwJsEDPUfyoyMnB4pWzjI/Gmh+M9+n4VIC7vlFC5JHOKaFOdvQetJtYt70ASnAYHrzQVHJJ70hAAGcn3pFOcigBmCp4NSDO7h8j0pJANuQee9NDE4AHNMYBcGnlQwI96NjZGAfenbXznbSEQEAZGaQ4xUjLhO/JpjjDcEEUwH7yOQOPSkxuPpSDG0hhhhSZZiABzQBKGbGG+lDE7cjgj9aesXy/MeTTWiYdDuH8qQDEbeNvepCoSiKMKxP5UrDc9NAREg9BmnIdvQDNOIUHHWgtjoD+VMB6sxPOP1odu1CdMkUxqADdTDtNMLE0mPWkUkXSqt1AP1pmFB+UAUhfJwKTODjvQIXdTweKhznpTkoAk7cUxVwfmySaXNKDng9aAEKgDNNGc9jSscHpShhTEKxwMVBI3b1p7uOT6VD1OTSY0AFLRRSKLO1QPuimugbp1p3rUbMQoIpkkeMN1NKr4OM00nNM6HNAFo8jPY0nP4jp71Gj4+lP3A0DHghhwaZnPBHI60jZHIpquS49e/vQIY1ApZBhz6UlDGhaSiikM//2Q=="
}

> convert -gaussian-blur 5x3 img/tiny-cy.jpg img/tiny-cy-blur.jpg
```

| Original | Tiny | Tiny Blurred |
| ![](./img/cy.jpg) | | |

## Debugging Tips
- Use `cmp` to determine if the first n bytes of a file are identical.
- The -o flag can be used to output the entire image for debuggling.

## References
Some discussion of this technique can be found online:
- https://stackoverflow.com/questioins/56236805/create-jpeg-thumb-image-with-general-fixed-header
- https://engineering.fb.com/android/the-technology-behind-preview-photos

![jpeg-reference](./img/jpeg-reference.jpg)

The full JPEG specification: https://www.w3.org/Graphics/JPEG/jfif3.pdf
