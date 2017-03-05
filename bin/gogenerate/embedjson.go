package main

import (
	"io"
	"os"
)

// Reads all .txt files in the current folder
// and encodes them as strings literals in textfiles.go
func main() {
	outFileName := "rules.go"
	out, _ := os.Create(outFileName)
	out.Write([]byte("package diffence\n\nconst (\n"))
	out.Write([]byte("\tdefaultRulesJSON = `"))

	inFileName := "test/fixtures/rules/gitrob.json"
	in, _ := os.Open(inFileName)

	io.Copy(out, in)
	out.Write([]byte("`\n"))
	out.Write([]byte(")\n"))
}
