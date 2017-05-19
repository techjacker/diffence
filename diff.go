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

// GetHashKey returns the hash key identifier for the diff
func (d *DiffItem) GetHashKey() string {
	if d.commit != "" {
		return fmt.Sprintf("%s:%s", d.commit, d.fPath)
	}
	return d.fPath
}

// SplitHashKey splits a DiffItem's hash key
func SplitDiffHashKey(s string) (string, string) {
	parts := strings.Split(s, ":")
	if len(parts) > 1 {
		return parts[0], parts[1]
	}
	return "", s
}

// Diff is a list of split diffs
type Diff struct {
	ignorer   Matcher
	Items     []DiffItem
	Error     error
	commitTmp string
}

// Push a diff on to the list
func (d *Diff) Push(s string) {

	var commitHeader, commit string

	if beginsWithHash(s) {
		commitHeader, s = split(s, "\n")
		commit = extractHash(commitHeader)
		d.commitTmp = commit
	}
	// add commit to diffs within each diff which do not
	// have the commit on the line directly above them
	if d.commitTmp != "" && commit == "" {
		commit = d.commitTmp
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
		commit: commit,
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

func extractHash(in string) string {
	return strings.Fields(in)[0]
}
