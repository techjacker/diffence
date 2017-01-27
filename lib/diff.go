package diffence

import "strings"

// DiffItem is a diff struct for an inidividual file
type DiffItem struct {
	raw string
	// filename  string
	// addedText string
	// match        bool
	// matchedRules []rule
}

// func (d *DiffItem) getHeader() []byte {
func (d *DiffItem) getHeader() string {
	newLineIndex := strings.Index(d.raw, "\n")
	return d.raw[:newLineIndex]
}

func (d *DiffItem) getFilename() string {
	prefix := "b/"
	pathBIndex := strings.Index(d.raw, prefix)
	newLineIndex := strings.Index(d.raw, "\n")
	return d.raw[pathBIndex+len(prefix) : newLineIndex]
}
