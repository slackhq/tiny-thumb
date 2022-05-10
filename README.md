// tiny_thumb takes an input image and produces a payload that can be used to
// reconstitute a small jpeg preview of the input.
//
// The basic idea is to convert the image to jpeg, scale it down, reduce it's
// 'quality' and use enough static parameters in the jpeg header such that the
// header can be stripped and reused across images.
//
// Some discussion of this technique can be found online:
// -
https://stackoverflow.com/questions/56236805/create-jpeg-thumb-image-with-general-fixed-header
// - https://engineering.fb.com/android/the-technology-behind-preview-photos
//
// There two options that can be changed; the 'type' and the maximum dimension.
// Each type corresponds to a specific jpeg header and dimension offset.
//
// The output of this program is a json object containing the key "Payload"
// whose value is a base64 encoded byte array. The base64 header can be found
// in the key "Debug.Head", and the dimension offset in Debug.DimensionOffset.
//
// It is expected that a server and client will preshare the mapping from all
// known types to there corresponding headers and dimension offsets. Upon
// receiving a payload a client can reconstitute the full JPEG by using the
// following process:
//
// - Split payload into three parts; the first byte is $type, the next four
// bytes are the $dimensions, and the remaining bytes is $tail.
// - Get the corresponding $header and $dimension_offset for $type. If you
// do not have values for this $type, fail.
// - Set the four bytes of $header beginning at offset $dimension_offset to
// the value of $dimensions.
// - Concatenate $header and $tail to get the final result.
//
// Useful tricks if debuggling this program:
// - Use `cmp` to determine if the first n bytes of a file are identical.
// - The -o flag can be used to output the entire image for debuggling.
// A good jpeg reference:
// -
http://vip.sugovica.hu/Sardi/kepnezo/JPEG%20File%20Layout%20and%20Format.htm
//
// WARNING: Changing type parameters like quality causes the header to
// change. This may break images you have already served via this program.
