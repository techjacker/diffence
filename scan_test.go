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
		want []DiffItem
	}{
		{
			name: "SplitDiffs()",
			args: args{r: getFixtureFile("test/fixtures/diffs/single.diff")},
			want: []DiffItem{
				{fPath: "README.md"},
			},
		},
		{
			name: "SplitDiffs()",
			args: args{r: getFixtureFile("test/fixtures/diffs/multi.diff")},
			want: []DiffItem{
				{fPath: "TODO.md"},
				{fPath: "systemdlogger/aws.py"},
				{fPath: "systemdlogger/cloudwatch.py"},
				{fPath: "tests/fixtures/config.json"},
				{fPath: "tests/test_aws.py"},
				{fPath: "tests/test_cloudwatch.py"},
				{fPath: "tests/test_runner_integration.py"},
			},
		},
		{
			name: "SplitDiffs()",
			args: args{r: getFixtureFile("test/fixtures/diffs/logp.truncated.diff")},
			want: []DiffItem{
				{fPath: "README.md"},
				{fPath: "TODO.md"},
				{fPath: "check.go"},
				{fPath: "results.go"},
			},
		},
		{
			name: "SplitDiffs()",
			args: args{r: getFixtureFile("test/fixtures/diffs/longline.diff")},
			want: []DiffItem{
				{fPath: "web/src/main/resources/static/js/menu.js"},
				{fPath: "web/src/main/resources/static/js/vendor/jquery-1.11.1.min.js"},
				{fPath: "web/src/main/resources/templates/layout.vm"},
			},
		},
		{
			name: "SplitDiffs()",
			args: args{r: getFixtureFile("test/fixtures/diffs/logp.diff")},
			want: generateWantDiffFromFiles("test/fixtures/diffs/expected/logp.diff.filepaths.txt"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Diff{}
			err := SplitDiffs(tt.args.r, &d)
			if err != nil {
				t.Logf("SplitDiffs =%d", len(d.Items))
				t.Fatalf("SplitDiffs threw error %#v", err)
			}
			for i, di := range d.Items {
				equals(t, tt.want[i].fPath, di.fPath)
			}
		})
	}
}
