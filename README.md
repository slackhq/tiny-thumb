# Tiny Thumb
Takes an image as input and produces a tiny payload that can be used to reconstitute a thumbnail of the input using preshared, static, image quality parameters. This method of compressing images is particularly useful in client/server models where payload size is a concern and custom client side manipulation of the payload, prior to display, is possible.

At the time of writing Slack uses this algorithm in its product to efficiently inline image previews into certain API responses. These previews are displayed to end users while the high resolution content is being fetched over the network.

## Example Usage
In the example below we will shrink a 13298 byte image down to 823 bytes and then apply a gaussian blur.

```
% du -b img/cy.jpg
13298   img/cy.jpg
% go run main.go -d 128 -o img/tiny-cy.jpg img/cy.jpg          
{
  "Debug": {
    "Parameters": {
      "Quality": 7,
      "Head": "/9j/2wCEAHJPVmRWR3JkXWSBeXKIq/+6q52dq//6/8////////////////////////////////////////////////////8BeYGBq5ar/7q6///////////////////////////////////////////////////////////////////////////AABEIAAAAAAMBIgACEQEDEQH/xAGiAAABBQEBAQEBAQAAAAAAAAAAAQIDBAUGBwgJCgsQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+gEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoLEQACAQIEBAMEBwUEBAABAncAAQIDEQQFITEGEkFRB2FxEyIygQgUQpGhscEJIzNS8BVictEKFiQ04SXxFxgZGiYnKCkqNTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqCg4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2dri4+Tl5ufo6ery8/T19vf4+fr/2gAMAwEAAhEDEQA/AA==",
      "DimensionOffset": 141
    },
    "Final": "/9j/2wCEAHJPVmRWR3JkXWSBeXKIq/+6q52dq//6/8////////////////////////////////////////////////////8BeYGBq5ar/7q6///////////////////////////////////////////////////////////////////////////AABEIAIAAXwMBIgACEQEDEQH/xAGiAAABBQEBAQEBAQAAAAAAAAAAAQIDBAUGBwgJCgsQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+gEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoLEQACAQIEBAMEBwUEBAABAncAAQIDEQQFITEGEkFRB2FxEyIygQgUQpGhscEJIzNS8BVictEKFiQ04SXxFxgZGiYnKCkqNTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqCg4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2dri4+Tl5ufo6ery8/T19vf4+fr/2gAMAwEAAhEDEQA/AIww5LE5pM9DSU4Fc5C8fWkMG5OR0perYwB603J65oxnpQK4pbOADnFDHvtFJ0bNDkHkUhjSDjnpTskrim0pGAOlMQq9aXJyfX3poJB4pzfdByaADPFJ9DRjIyOlHOM8UAIf1pV47Zpe33cUbTt7YoAQkE98UqgE9TR17DHrSYyeaAFYH8qaBwacOuCaUjkjpQBHR160UUAOUe340rf5zSYO3NPbhQfzoAjpDTs5HSmmgQ8g574pRtzhenrTWJxTwBtDce9Ax2AvHUdaQYbdjFO4xgdKaflXHbtSAGAOO4puzkAHrQOXGKczfdJFACeTg5zxSAYORwKeD1+lNcBVHT60AJx1GcntQxyM9KQD5h1NOIPUHj0pjGZJHfAppoyRRQIsKFKg4zTeEAxyD60q5A7ZpQM49qQBxjGaXIxjtTWH7wE96CN3SgA2jqO3vQ4yAOtKwUIcHmkRwFAI5oAToee9Mc5AoOXY44FPMRC5DfhTAjHDDBIGaUvgkA59KcmCoBAIHao2yPYdqAF5xyPpTDTx8wy2aYetMCfBAz1H8qMjJx0pWBxkfjTQ/Ge/SpAXd8oNCZJHOKaFOSvQetIFYt3zQBLwGHvQVGM55pCAMdT70inIIoAYAVPWpBnON+VpsoHUde9IGLEBRTGKq4YClZQw47Uu1twwpxTirDPy9aQiAgDjNIcGpCuAoOcd6jPBPIP0pgSbyOQOPSkxuPpSDG0hhhhSZZmAHWgCUM2MNQ5O0HoR+tPWL5fmPJ9Ka0TDvkfypANQ7xt708gJSxIFYmhhuc00BESD0Gaeh29AM0p259aC2OgP5UwHqzE8/wBaR25xSr93JFRv60AKWphC1GSWpcUiki2VVuoB+tNwoPygCkL5OBTc4OO9Ah26n5qHOelOSgB56cdaaq4PzZJpc0oOevWgBGUAZ5FNXJPQGlJwcEUoYdqYgY9qhkPanu4wTUI9TSY0LRRRSKLO1QPuimugbp1p3rUbsQoIpkkQGCeTT1fBphOaZ0OaALR9expMn8R+tRo+PpT8g0AOzuXg0zcCOhBHX2obI5FIjFn/AJigCJutKKGGHIooY0LSUUUhn//Z",
    "Height": 128,
    "Width": 95,
    "PayloadLen": 823,
    "MaxDimension": 128
  },
  "Payload": "AQCAAF+MMOSxOaTPQ0lOBXOQvH1pDBuTkdKXq2MAetNyeuaMZ6UCuKWzgA5xQx77RSdGzQ5B5FIY0g456U7JK4ptKRgDpTEKvWlycn196aCQeKc33QcmgAzxSfQ0YyMjpRzjPFACH9aVeO2aXt93FG07e2KAEJBPfFKoBPU0dewx60mMnmgBWB/KmgcGnDrgmlI5I6UAR0detFFADlHt+NK3+c0mDtzT24UH86AI6Q07OR0ppoEPIOe+KUbc4Xp601icU8AbQ3HvQMdgLx1HWkGG3YxTuMYHSmn5Vx27UgBgDjuKbs5AB60DlxinM33SRQAnk4Oc8UgGDkcCng9fpTXAVR0+tACcdRnJ7UMcjPSkA+YdTTiD1B49KYxmSR3wKaaMkUUCLChSoOM03hAMcg+tKuQO2aUDOPakAcYxmlyMY7U1h+8BPegjd0oANo6jt70OMgDrSsFCHB5pEcBQCOaAE6HnvTHOQKDl2OOBTzEQuQ34UwIxwwwSBmlL4JAOfSnJgqAQCB2qNsj2HagBeccj6Uw08fMMtmmHrTAnwQM9R/KjIycdKVgcZH400Pxnv0qQF3fKDQmSRzimhTkr0HrSBWLd80AS8Bh70FRjOeaQgDHU+9IpyCKAGAFT1qQZzjflabKB1HXvSBixAUUxiquGApWUMOO1LtbcMKcU4qwz8vWkIgIA4zSHBqQrgKDnHeozwTyD9KYEm8jkDj0pMbj6UgxtIYYYUmWZgB1oAlDNjDUOTtB6EfrT1i+X5jyfSmtEw75H8qQDUO8be9PICUsSBWJoYbnNNAREg9BmnodvQDNKdufWgtjoD+VMB6sxPP8AWkducUq/dyRUb+tAClqYQtRklqXFIpItlVbqAfrTcKD8oApC+TgU3ODjvQIdup+ahznpTkoAeenHWmquD82SaXNKDnr1oARlAGeRTVyT0BpScHBFKGHamIGPaoZD2p7uME1CPU0mNC0UUUiiztUD7oproG6dad61G7EKCKZJEBgnk09XwaYTmmdDmgC0fXsaTJ/EfrUaPj6U/INADs7l4NM3AjoQR19qGyORSIxZ/wCYoAibrSihhhyKKGNC0lFFIZ//2Q=="
}

% convert -gaussian-blur 5x3 img/tiny-cy.jpg img/tiny-cy-blur.jpg
```

| Original | Tiny | Tiny Blurred |
| --- | --- | --- |
| ![](./img/cy.jpg) | ![](./img/tiny-cy.jpg) | ![](./img/tiny-cy-blur.jpg) |

## High Level Tiny Thumb Generation Algorithim
- Convert the image to a jpeg.
- Scale it down.
- Reduce its 'quality' using predetermined, hardcoded parameters.
- Strip the jpeg header.

## Program Output and Reconstitution Algorithim
The output of this program is a json object containing the key `Payload` whose value is a base64 encoded byte array. The base64 header can be found in the key `Debug.Head`, and the dimension offset in `Debug.DimensionOffset`.

It is expected that a server and client will preshare the mapping from all known types to there corresponding headers and dimension offsets. Upon receiving a payload a client can reconstitute the full JPEG by using the following process:

- Split payload into three parts; the first byte is the `$type`, the next four bytes are the `$dimensions`, and the remaining bytes are the `$tail`.
- Get the corresponding preshared `$header` and `$dimension_offset` for `$type`. If you do not have values for this `$type`, fail.
- Set the four bytes of `$header` beginning at offset `$dimension_offset` to the value of `$dimensions`.
- Concatenate `$header` and `$tail` to get the final result.
- Optional: apply a gaussian blur. :sparkles:

## Options and Configuration
This program takes two options that can be changed; the 'type' and the 'maximum dimension'. Each type corresponds to a specific jpeg 'header' and 'dimension offset'. This program includes a mapping of types currently used by Slack, the details of which can be found in the program output or source code.

## Debugging Tips
- Use `cmp` to determine if the first n bytes of a file are identical.
- The -o flag can be used to output the entire image for debuggling.

## References
![jpeg-reference](./img/jpeg-reference.jpg)

The full JPEG specification: https://www.w3.org/Graphics/JPEG/jfif3.pdf

Some discussion of this technique can be found online:
- https://stackoverflow.com/questioins/56236805/create-jpeg-thumb-image-with-general-fixed-header
- https://engineering.fb.com/android/the-technology-behind-preview-photos
