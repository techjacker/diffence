package diffence

import (
	"bytes"
	"encoding/hex"
	"unicode"

	"github.com/y0ssar1an/q"
)

const (
	diffSep = "diff --git a"
)

// ScanDiffs splits on the diff of an inidividual file
func ScanDiffs(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	k, newLineIndex, prevNewLineIndex, diffSepEndIndex := 0, 0, 0, 0
	dataLen := len(data) - 1
	diffSepLen := len(diffSep)

	// loop until no more bytes left to read in this chunk of data
	for k < dataLen {
		// find the next newline
		if i := bytes.IndexByte(data[k:], '\n'); i >= 0 {
			// how far advanced already (k)
			// index after last \n char (+ i)
			// start at next byte (+ 1)
			newLineIndex = k + i + 1
			diffSepEndIndex = newLineIndex + diffSepLen
			if diffSepEndIndex < dataLen && string(data[newLineIndex:diffSepEndIndex]) == diffSep {
				if prevNewLineIndex > 0 {

					hashEnd := bytes.IndexFunc(data[prevNewLineIndex:newLineIndex], unicode.IsSpace)
					commitHash := string(data[prevNewLineIndex : prevNewLineIndex+hashEnd])
					_, e := hex.DecodeString(commitHash)
					if e == nil {
						// q.Q(commitHash)
						q.Q(prevNewLineIndex)
						// q.Q(dataLen)
						tokenEnd := prevNewLineIndex - 1
						if tokenEnd < 0 {
							tokenEnd = 0
						}
						return prevNewLineIndex, dropCR(data[0:tokenEnd]), nil
						// return prevNewLineIndex, dropCR(data[0 : k+i]), nil
					}
				}
				return newLineIndex, dropCR(data[0 : k+i]), nil
			}
			prevNewLineIndex = newLineIndex
			// prevNewLineIndex = newLineIndex - 1
			k += i + 1
		} else {
			k = dataLen
		}
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
