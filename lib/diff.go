package diffence

import (
	"bufio"
	"io"
)

type diffItem struct {
	data []byte
	// filename     []byte
	// addedText    []byte
	// match        bool
	// matchedRules []rule
}

// Differ creates diffItems from a raw git diff text input
type Differ interface {
	// split(string) []string
	Parse(io.Reader) []diffItem
}

// NewDiffer is a Differ factory
func NewDiffer() Differ {
	return &diff{}
}

type diff struct{}

// Parse splits a diff into individual file diffs and parses each one
// in a separate go routine
func (d diff) Parse(r io.Reader) []diffItem {

	// Default scanner is bufio.ScanLines. Lets use ScanWords.
	// Could also use a custom function of SplitFunc type
	scanner := bufio.NewScanner(r)
	scanner.Split(ScanDiffs)

	items := []diffItem{}
	// i := 0
	for scanner.Scan() {
		word := scanner.Bytes()
		// fmt.Printf("%d: %#s\n", i, pretty.Formatter(word))
		// i += 1
		items = append(items, diffItem{word})
	}

	return items
}
