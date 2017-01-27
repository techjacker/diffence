package main

import (
	"bytes"

	df "github.com/techjacker/diffence/lib"
)

func main() {
	r := bytes.NewReader([]byte("hello world"))
	items, err := df.SplitDiffs(r)
	println("items", items)
	println("err", err)
}
