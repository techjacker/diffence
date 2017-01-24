package diffence

import (
	"bufio"
	"bytes"
	"io"
	"reflect"
	"testing"
)

// func TestScanDiffs(t *testing.T) {
// 	type args struct {
// 		data  []byte
// 		atEOF bool
// 	}
// 	tests := []struct {
// 		name        string
// 		args        args
// 		wantAdvance int
// 		wantToken   []byte
// 		wantErr     bool
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gotAdvance, gotToken, err := ScanDiffs(tt.args.data, tt.args.atEOF)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("ScanDiffs() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if gotAdvance != tt.wantAdvance {
// 				t.Errorf("ScanDiffs() gotAdvance = %v, want %v", gotAdvance, tt.wantAdvance)
// 			}
// 			if !reflect.DeepEqual(gotToken, tt.wantToken) {
// 				t.Errorf("ScanDiffs() gotToken = %v, want %v", gotToken, tt.wantToken)
// 			}
// 		})
// 	}
// }

func TestScanDiffsWithBufioScanner(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "Differ.Parse()",
			args: args{r: bytes.NewReader(
				[]byte(
					"diff --git a/README.md b/README.md" +
						"\n" +
						"index 82366e3..5fc99b9 100644" +
						"\n" +
						"diff --git a/TODO.md b/TODO.md" +
						"\n" +
						"index 82366e3..5fc99b9 100644",
				),
			)},
			want: [][]byte{
				[]byte(
					"diff --git a/README.md b/README.md" +
						"\n" +
						"index 82366e3..5fc99b9 100644",
				),
				[]byte(
					"diff --git a/TODO.md b/TODO.md" +
						"\n" +
						"index 82366e3..5fc99b9 100644",
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := [][]byte{}
			scanner := bufio.NewScanner(tt.args.r)
			scanner.Split(ScanDiffs)
			for scanner.Scan() {
				got = append(got, scanner.Bytes())
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diff.Parse() \n\nGOT: %s, \n\nWANT: %s", got, tt.want)
			}
		})
	}
}
