# Tiny Thumb
Takes an input image and produces a payload that can be used to reconstitute a tiny preview of the input using preshared, static, quality parameters. This method is particularly useful in client/server models where parameters can be preshared, jpeg support is available on teh client, and payload size is a concern.

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

## Debugging Tips
- Use `cmp` to determine if the first n bytes of a file are identical.
- The -o flag can be used to output the entire image for debuggling.

## References
Some discussion of this technique can be found online:
- https://stackoverflow.com/questioins/56236805/create-jpeg-thumb-image-with-general-fixed-header
- https://engineering.fb.com/android/the-technology-behind-preview-photos

The full JPEG specification: https://www.w3.org/Graphics/JPEG/jfif3.pdf
