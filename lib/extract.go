package diffence

import (
	"bytes"
	"unicode"
)

func extractFileName(in []byte) []byte {
	// prefix := []byte("diff --git a/")
	// firstLine := bytes.SplitN(in, []byte)
	newLineIndex := bytes.IndexByte(in, '\n')
	fileNameWithPrefixIndex := bytes.LastIndexFunc(in[0:newLineIndex], unicode.IsSpace) + 1
	return bytes.TrimPrefix(in[fileNameWithPrefixIndex:newLineIndex], []byte("b/"))
}

func extractAddedText(in []byte) []byte {
	return []byte("")
}
