package main

import (
	"flag"
	"fmt"
	"net/url"
)

func main() {
	s := flag.String("s", "", "String to encode/decode")
	d := flag.Bool("d", false, "decode")
	de := flag.Bool("de", false, "double encode")

	flag.Parse()

	output := ""

	if *d {
		decoded, err := url.QueryUnescape(*s)
		if err != nil {
			panic(err)
		}
		output = decoded
	} else if *de {
		output = url.QueryEscape(url.QueryEscape(*s))
	} else {
		output = url.QueryEscape(*s)
	}

	fmt.Println(output)
}
