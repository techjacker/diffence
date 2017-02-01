package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/kr/pretty"
	df "github.com/techjacker/diffence/lib"
)

func main() {
	r := bytes.NewReader([]byte("hello world"))

	rules, err := df.ReadRulesFromFile("test/fixtures/rules/gitrob.json")
	if err != nil {
		fmt.Printf("\nCannot read rule file: %s\n", err)
		os.Exit(1)
		return
	}

	res, err := df.CheckDiffs(r, rules)
	if err != nil {
		fmt.Printf("\nError reading diff\n%s\n", err)
		os.Exit(1)
		return
	}
	fmt.Printf("%#v\n", pretty.Formatter(res))
}
