package diffence

import (
	"bufio"
	"bytes"
	"io"
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

func TestSplitDiffs(t *testing.T) {
	type args struct {
		r io.Reader
	}

	tests := []struct {
		name string
		args args
		want []wantDiff
	}{
		{
			name: "Differ.Parse()",
			args: args{
				r: getFixtureFile("test/fixtures/diffs/single.diff"),
			},
			want: []wantDiff{
				wantDiff{
					header:   "diff --git a/README.md b/README.md",
					filename: "README.md",
				},
			},
		},
		{
			name: "Differ.Parse()",
			args: args{
				r: getFixtureFile("test/fixtures/diffs/multi.diff"),
			},
			want: []wantDiff{
				wantDiff{
					header:   "diff --git a/TODO.md b/TODO.md",
					filename: "TODO.md",
				},
				wantDiff{
					header:   "diff --git a/systemdlogger/aws.py b/systemdlogger/aws.py",
					filename: "systemdlogger/aws.py",
				},
				wantDiff{
					header:   "diff --git a/systemdlogger/cloudwatch.py b/systemdlogger/cloudwatch.py",
					filename: "systemdlogger/cloudwatch.py",
				},
				wantDiff{
					header:   "diff --git a/tests/fixtures/config.json b/tests/fixtures/config.json",
					filename: "tests/fixtures/config.json",
				},
				wantDiff{
					header:   "diff --git a/tests/test_aws.py b/tests/test_aws.py",
					filename: "tests/test_aws.py",
				},
				wantDiff{
					header:   "diff --git a/tests/test_cloudwatch.py b/tests/test_cloudwatch.py",
					filename: "tests/test_cloudwatch.py",
				},
				wantDiff{
					header:   "diff --git a/tests/test_runner_integration.py b/tests/test_runner_integration.py",
					filename: "tests/test_runner_integration.py",
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
				equals(t, extractHeader(di.raw), tt.want[i].header)
				equals(t, di.filePath, tt.want[i].filename)
			}
		})
	}
}

type wantDiff struct {
	header   string
	filename string
}

func Test_extract(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want wantDiff
	}{
		{
			name: "diff.getHeader()",
			args: args{
				in: "diff --git a/README.md b/README.md" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want: wantDiff{
				header:   "diff --git a/README.md b/README.md",
				filename: "README.md",
			},
		},
		{
			name: "diff.getHeader()",
			args: args{
				in: "diff --git a/TODO.md b/TODO.md" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want: wantDiff{
				header:   "diff --git a/TODO.md b/TODO.md",
				filename: "TODO.md",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			equals(t, extractHeader(tt.args.in), tt.want.header)
			equals(t, extractFilePath(tt.args.in), tt.want.filename)
		})
	}
}
