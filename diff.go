package diffence

import (
	"encoding/hex"
	"fmt"
	"strings"
	"unicode"
)

const (
	pathSep = " b/"
)

// List is an interface for adding items to a list
type List interface {
	Push(string)
}

// DiffItem is a diff struct for an inidividual file
type DiffItem struct {
	raw    string
	fPath  string
	commit string
}

// Diff is a list of split diffs
type Diff struct {
	ignorer         Matcher
	Items           []DiffItem
	Error           error
	commitHeaderTmp string
}

// Push a diff on to the list
func (d *Diff) Push(s string) {

	var commitHeader string

	if beginsWithHash(s) {
		commitHeader, s = split(s, "\n")
		d.commitHeaderTmp = commitHeader
	}
	// add commitHeader to diffs within each diff which do not
	// have the commitHeader on the line directly above them
	if d.commitHeaderTmp != "" && commitHeader == "" {
		commitHeader = d.commitHeaderTmp
	}

	fPath, err := extractFilePath(s)
	if err != nil {
		d.Error = err
		return
	}

	if d.ignorer != nil && d.ignorer.Match(fPath) {
		return
	}

	d.Items = append(d.Items, DiffItem{
		raw:    s,
		fPath:  fPath,
		commit: commitHeader,
	})
}

// split out logic from scan.go
func beginsWithHash(s string) bool {
	if len(s) < 1 {
		return false
	}
	hashEnd := strings.IndexFunc(s, unicode.IsSpace)
	if hashEnd < 1 {
		return false
	}
	commitHash := s[:hashEnd]
	_, e := hex.DecodeString(commitHash)
	if e == nil {
		return true
	}
	return false
}

/////////////////////////////////////////
/////////////////////////////////////////
/////////////////////////////////////////
func split(s, sep string) (string, string) {
	// Empty string should just return empty
	if len(s) == 0 {
		return s, s
	}
	slice := strings.SplitN(s, sep, 2)
	// Incase no separator was present
	if len(slice) == 1 {
		return slice[0], ""
	}
	return slice[0], slice[1]
}

func extractFilePath(in string) (string, error) {
	pathBIndex := strings.Index(in, pathSep)
	newLineIndex := strings.Index(in, "\n")
	if pathBIndex >= 0 && newLineIndex > pathBIndex {
		return in[pathBIndex+len(pathSep) : newLineIndex], nil
	}
	return "", fmt.Errorf("Not valid diff content:\n%s", in)
}
