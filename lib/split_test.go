package diffence

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"path"
	"reflect"
	"testing"
)

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
			name: "ScanDiffs() split fn",
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

func getFixtureFile(filename string) io.Reader {
	cwd, _ := os.Getwd()
	file, err := os.Open(path.Join(cwd, "../", filename))
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func Test_diff_Parse(t *testing.T) {
	type args struct {
		r io.Reader
	}
	type want struct {
		header   string
		filename []byte
	}
	tests := []struct {
		name string
		args args
		want []want
	}{
		{
			name: "Differ.Parse()",
			args: args{
				r: getFixtureFile("test/fixtures/diffs/single.diff"),
			},
			want: []want{
				want{
					header:   "diff --git a/README.md b/README.md",
					filename: []byte("README.md"),
				},
			},
		},
		{
			name: "Differ.Parse()",
			args: args{
				r: getFixtureFile("test/fixtures/diffs/multi.diff"),
			},
			want: []want{
				want{
					header:   "diff --git a/TODO.md b/TODO.md",
					filename: []byte("TODO.md"),
				},
				want{
					header:   "diff --git a/systemdlogger/aws.py b/systemdlogger/aws.py",
					filename: []byte("systemdlogger/aws.py"),
				},
				want{
					header:   "diff --git a/systemdlogger/cloudwatch.py b/systemdlogger/cloudwatch.py",
					filename: []byte("systemdlogger/cloudwatch.py"),
				},
				want{
					header:   "diff --git a/tests/fixtures/config.json b/tests/fixtures/config.json",
					filename: []byte("tests/fixtures/config.json"),
				},
				want{
					header:   "diff --git a/tests/test_aws.py b/tests/test_aws.py",
					filename: []byte("tests/test_aws.py"),
				},
				want{
					header:   "diff --git a/tests/test_cloudwatch.py b/tests/test_cloudwatch.py",
					filename: []byte("tests/test_cloudwatch.py"),
				},
				want{
					header:   "diff --git a/tests/test_runner_integration.py b/tests/test_runner_integration.py",
					filename: []byte("tests/test_runner_integration.py"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// check for error scanning
			items, err := SplitDiffs(tt.args.r)
			if err != nil {
				t.Fatalf("SplitDiffs threw error %#v", err)
			}

			// check extracting metadata
			for i, di := range items {
				if tt.want[i].header != di.getHeader() {
					t.Errorf("SplitDiffs() item:%d \nWANT: %s\nGOT: %s", i, tt.want[i].header, di.getHeader())
					t.Fatalf("Body:\n\n%s", di.raw)
				}
			}
		})
	}
}
