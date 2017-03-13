package diffence

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path/filepath"
)

// Matcher is an interface for matching against string inputs
type Matcher interface {
	Match(string) bool
}

// NewIgnorerFromFile will safely create an Ignorer whether the file exists or not
func NewIgnorerFromFile(fPath string) *Ignorer {
	return &Ignorer{getFileAsSlice(fPath)}
}

// NewIgnorer will create an Ignorer from a read stream
func NewIgnorer(r io.Reader) *Ignorer {
	return &Ignorer{splitLines(r)}
}

// Ignorer is used to exclude content in .secignore files
type Ignorer struct {
	patterns []string
}

// Match reports whether a filepath is listed in Ignorer.patterns[]string
func (i Ignorer) Match(in string) bool {
	for _, p := range i.patterns {
		// http://golang-jp.org/pkg/path/filepath/#Match
		if matched, _ := filepath.Match(p, in); matched == true {
			return true
		}
	}
	return false
}

// getFileAsSlice returns the contents of a file as a slice of strings
func getFileAsSlice(fPath string) []string {
	return splitLines(getFile(fPath))
}

func getFile(fPath string) io.Reader {
	if _, err := os.Stat(fPath); os.IsNotExist(err) {
		return bytes.NewBuffer([]byte{})
	}
	f, err := os.Open(fPath) // /path/to/whatever exists
	if err != nil {
		return bytes.NewBuffer([]byte{})
	}
	return f
}

// Split
func splitLines(r io.Reader) []string {
	if v, ok := r.(io.ReadCloser); ok == true {
		defer v.Close()
	}
	l := []string{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		l = append(l, scanner.Text())
	}
	return l
}
