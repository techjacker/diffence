package diffence

import "bytes"

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

// ScanDiffs splits on the diff of an inidividual file
func ScanDiffs(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	k, dataLen := 0, len(data)-1
	for k < dataLen {
		if i := bytes.IndexByte(data[k:], '\n'); i >= 0 {
			if k+i+1 < dataLen && string(data[k+i+1]) == "d" {
				return k + i + 1, dropCR(data[0 : k+i]), nil
				// start at index after last \n char
			}
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
