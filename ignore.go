package diffence

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path/filepath"
)

type Matcher interface {
	Match(string) bool
}

type Ignorer struct {
	patterns []string
}

func (i Ignorer) Match(in string) bool {
	for _, p := range i.patterns {
		// http://golang-jp.org/pkg/path/filepath/#Match
		if matched, _ := filepath.Match(p, in); matched == true {
			return true
		}
	}
	return false
}

func getFileContents(fPath string) []string {
	return SplitLines(getFile(fPath))
}

func getFile(fPath string) io.Reader {

	if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
		return bytes.NewBuffer([]byte{})
	}
	f, err := os.Open(fPath) // path/to/whatever exists
	if err != nil {
		return bytes.NewBuffer([]byte{})
	}
	return f
}

func SplitLines(r io.Reader) []string {
	l := []string{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		l = append(l, scanner.Text())
	}
	return l
}
