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

// func (d *DiffItem) getFilename() []byte {
// 	// prefix := []byte("diff --git a/")
// 	// firstLine := bytes.SplitN(in, []byte)
// 	newLineIndex := bytes.IndexByte(d.raw, '\n')
// 	indexfilePrefix := bytes.LastIndexFunc(
// 		d.raw[0:newLineIndex],
// 		unicode.IsSpace,
// 	)
// 	prefix := d.raw[indexfilePrefix+1 : newLineIndex]
// 	return bytes.TrimPrefix(prefix, []byte("b/"))
// }
