package diffence

import (
	"bufio"
	"io"
)

// DiffItem is a diff struct for an inidividual file
type DiffItem struct {
	raw       []byte
	filename  []byte
	addedText []byte
	// match        bool
	// matchedRules []rule
}

// NewDiffItem is a DiffItem factory
func NewDiffItem(raw []byte) DiffItem {
	return DiffItem{
		raw:       raw,
		filename:  extractFileName(raw),
		addedText: extractAddedText(raw),
	}
}

// Differ creates DiffItems from a raw git diff text input
type Differ interface {
	Parse(io.Reader)
}

// NewDiffer is a Differ factory
func NewDiffer() Differ {
	return &diff{}
}

type diff struct {
	items []DiffItem
}

// Parse splits a diff into individual file diffs and parses each one
// in a separate go routine
func (d *diff) Parse(r io.Reader) {

	scanner := bufio.NewScanner(r)
	scanner.Split(ScanDiffs)

	for scanner.Scan() {
		word := scanner.Bytes()
		d.items = append(d.items, NewDiffItem(word))
	}
}
