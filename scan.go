package diffence

import "bytes"

const (
	diffSep = "diff --git a"
)

func isDiffHeader(data *[]byte, newLineIndex int) bool {
	diffSepLen := len(diffSep)
	diffSepEndIndex := newLineIndex + diffSepLen
	dataLen := len(*data) - 1
	return diffSepEndIndex < dataLen && string((*data)[newLineIndex:diffSepEndIndex]) == diffSep
}

// ScanDiffs splits on the diff of an inidividual file
func ScanDiffs(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	k, newLineIndex := 0, 0
	dataLen := len(data) - 1

	// loop until no more bytes left to read in this chunk of data
	for k < dataLen {

		// find the next newline
		if i := bytes.IndexByte(data[k:], '\n'); i >= 0 {
			// k = index of scanned through data so far
			// i = index after last \n char
			// 1 = start at next byte
			newLineIndex = k + i + 1

			if beginsWithHash(string(data[newLineIndex:])) {
				return newLineIndex, dropCR(data[0 : newLineIndex-1]), nil
			}

			if isDiffHeader(&data, newLineIndex) {
				// if previous line does not begin with a hash then separate on diff headers
				if !beginsWithHash(string(data[0 : k+i])) {
					return newLineIndex, dropCR(data[0 : newLineIndex-1]), nil
				}
			}

			////////////////////////////////////////////////////////////////
			// keep advancing through data
			// k = index of scanned so far
			// i = index of new line
			// 1 = start at next byte after previous new line
			////////////////////////////////////////////////////////////////
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
