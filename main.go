// tiny_thumb takes an input image and produces a payload that can be used to
// reconstitute a small jpeg preview of the input.
//
// See the comment in pkg/tiny_thumb.go for more details.
//
// This program writes a json object to stdout containing the key "Payload"
// whose value is a base64 encoded byte array. The base64 header can be found
// in the key "Debug.Head", and the dimension offset in Debug.DimensionOffset.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	tt "github.com/slackhq/tiny-thumb/pkg"
)

var (
	typeID               = flag.Int("t", 1, "tiny thumb type id")
	dimen                = flag.Int("d", 32, "max dimension")
	out                  = flag.String("o", "", "debugging: if non-empty write the full result image to this path")
	checkParametersMatch = flag.Bool("checks", true, "tests that output matches known parameters")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: tiny_thumb file.jpg\n")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}
	b, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		panic(fmt.Errorf("error reading '%s': %v", flag.Arg(1), err))
	}
	r, err := tt.TinyThumb(b, byte(*typeID), *dimen, *checkParametersMatch)
	check(err)
	if *out != "" {
		err = ioutil.WriteFile(*out, r.Debug.Final, 0666)
		if err != nil {
			panic(fmt.Errorf("error writing: %v", err))
		}
	}
	jb, err := json.MarshalIndent(r, "", "  ")
	check(err)
	fmt.Println(string(jb))
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}
