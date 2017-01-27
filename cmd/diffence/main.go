package main

import (
	"bytes"

	df "github.com/techjacker/diffence/lib"
)

func main() {
	d := df.NewDiffer()
	d.Parse(bytes.NewReader([]byte("hello world")))
}
