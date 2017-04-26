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

func getTokenEnd(prevNewLineIndex int) int {
	tokenEnd := prevNewLineIndex - 1
	if tokenEnd >= 0 {
		return tokenEnd
	}
	return 0
}

// ScanDiffs splits on the diff of an inidividual file
func ScanDiffs(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// k, newLineIndex, prevNewLineIndex := 0, 0, 0
	k, newLineIndex := 0, 0
	dataLen := len(data) - 1

	// loop until no more bytes left to read in this chunk of data
	for k < dataLen {

		///////////////////////////
		// find the next newline //
		///////////////////////////
		if i := bytes.IndexByte(data[k:], '\n'); i >= 0 {
			// k = index of scanned through data so far
			// index after last \n char (+ i)
			// start at next byte (+ 1)
			newLineIndex = k + i + 1

			/////////////////////////////////
			// is this line a diff header? //
			/////////////////////////////////
			if isDiffHeader(&data, newLineIndex) {

				/////////////////////////////////////////////////////////
				// is the line before the diff header a commit header? //
				/////////////////////////////////////////////////////////
				// if prevNewLineIndex > 0 && beginsWithHash(string(data[prevNewLineIndex:newLineIndex])) {
				// 	// advance = ends = new line BELOW git diff header
				// 	// prev token = ends = new line of git diff header
				// 	return newLineIndex, dropCR(data[0:getTokenEnd(prevNewLineIndex)]), nil
				// }
				return newLineIndex, dropCR(data[0 : k+i]), nil
			}
			////////////////////////////////////////////////////////////////////
			// keep track of last new line - it could be the commit ID header //
			////////////////////////////////////////////////////////////////////
			// prevNewLineIndex = newLineIndex

			////////////////////////////////////////////////////////////////
			// keep advancing through data -> k = index of scanned so far //
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
