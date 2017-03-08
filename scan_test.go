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
		name     string
		args     args
		want     []wantDiff
		lenDiffs int
	}{
		{
			name: "SplitDiffs()",
			args: args{
				r: getFixtureFile("test/fixtures/diffs/single.diff"),
			},
			want: []wantDiff{
				{
					header:   "diff --git a/README.md b/README.md",
					filepath: "README.md",
				},
			},
		},
		{
			name: "SplitDiffs()",
			args: args{
				r: getFixtureFile("test/fixtures/diffs/multi.diff"),
			},
			want: []wantDiff{
				{
					header:   "diff --git a/TODO.md b/TODO.md",
					filepath: "TODO.md",
				},
				{
					header:   "diff --git a/systemdlogger/aws.py b/systemdlogger/aws.py",
					filepath: "systemdlogger/aws.py",
				},
				{
					header:   "diff --git a/systemdlogger/cloudwatch.py b/systemdlogger/cloudwatch.py",
					filepath: "systemdlogger/cloudwatch.py",
				},
				{
					header:   "diff --git a/tests/fixtures/config.json b/tests/fixtures/config.json",
					filepath: "tests/fixtures/config.json",
				},
				{
					header:   "diff --git a/tests/test_aws.py b/tests/test_aws.py",
					filepath: "tests/test_aws.py",
				},
				{
					header:   "diff --git a/tests/test_cloudwatch.py b/tests/test_cloudwatch.py",
					filepath: "tests/test_cloudwatch.py",
				},
				{
					header:   "diff --git a/tests/test_runner_integration.py b/tests/test_runner_integration.py",
					filepath: "tests/test_runner_integration.py",
				},
			},
		},
		{
			name: "SplitDiffs()",
			args: args{
				r: getFixtureFile("test/fixtures/diffs/logp.truncated.diff"),
			},
			want: []wantDiff{
				{
					header:   "diff --git a/README.md b/README.md",
					filepath: "README.md",
				},
				{
					header:   "diff --git a/TODO.md b/TODO.md",
					filepath: "TODO.md",
				},
				{
					header:   "diff --git a/check.go b/check.go",
					filepath: "check.go",
				},
				{
					header:   "diff --git a/results.go b/results.go",
					filepath: "results.go",
				},
			},
		},
		{
			name: "SplitDiffs()",
			args: args{
				r: getFixtureFile("test/fixtures/diffs/longline.diff"),
			},
			want: []wantDiff{
				{
					header:   "diff --git a/web/src/main/resources/static/js/menu.js b/web/src/main/resources/static/js/menu.js",
					filepath: "web/src/main/resources/static/js/menu.js",
				},
				{
					header:   "diff --git a/web/src/main/resources/static/js/vendor/jquery-1.11.1.min.js b/web/src/main/resources/static/js/vendor/jquery-1.11.1.min.js",
					filepath: "web/src/main/resources/static/js/vendor/jquery-1.11.1.min.js",
				},
				{
					header:   "diff --git a/web/src/main/resources/templates/layout.vm b/web/src/main/resources/templates/layout.vm",
					filepath: "web/src/main/resources/templates/layout.vm",
				},
			},
		},
		{
			name: "SplitDiffs()",
			args: args{
				r: getFixtureFile("test/fixtures/diffs/logp.diff"),
			},
			want: generateWantDiffFromFiles(
				"test/fixtures/diffs/expected/logp.diff.headers.txt",
				"test/fixtures/diffs/expected/logp.diff.filepaths.txt",
			),
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
				header, _ := extractHeader(di.raw)
				equals(t, tt.want[i].header, header)
				equals(t, tt.want[i].filepath, di.filePath)
			}
		})
	}
}

type wantDiff struct {
	header   string
	filepath string
}

type wantErr struct {
	header   bool
	filepath bool
}

func TestExtract(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    wantDiff
		wantErr wantErr
	}{

		{
			name: "diff.getHeader() - Admin Password",
			args: args{
				in: "diff --git a/web/src/main/resources/db/migration/V0_4__AdminPassword.sql b/web/src/main/resources/db/migration/V0_4__AdminPassword.sql" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want: wantDiff{
				header:   "diff --git a/web/src/main/resources/db/migration/V0_4__AdminPassword.sql b/web/src/main/resources/db/migration/V0_4__AdminPassword.sql",
				filepath: "web/src/main/resources/db/migration/V0_4__AdminPassword.sql",
			},
			wantErr: wantErr{false, false},
		},
		{
			name: "diff.getHeader()",
			args: args{
				in: "diff --git a/lib/check.go b/lib/check.go" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want: wantDiff{
				header:   "diff --git a/lib/check.go b/lib/check.go",
				filepath: "lib/check.go",
			},
			wantErr: wantErr{false, false},
		},
		{
			name: "diff.getHeader()",
			args: args{
				in: "diff --git a/README.md b/README.md" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want: wantDiff{
				header:   "diff --git a/README.md b/README.md",
				filepath: "README.md",
			},
			wantErr: wantErr{false, false},
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
				filepath: "TODO.md",
			},
			wantErr: wantErr{false, false},
		},
		{
			name: "diff.getHeader()",
			args: args{
				in: "hello world",
			},
			want: wantDiff{
				header:   "",
				filepath: "",
			},
			wantErr: wantErr{true, true},
		},
		{
			name: "diff.getHeader()",
			args: args{
				in: "diff --git a/cmd/diffence/main.go b/cmd/diffence/main.go" +
					"\n" +
					"index 08044af..098342c 100644",
			},
			want: wantDiff{
				header:   "diff --git a/cmd/diffence/main.go b/cmd/diffence/main.go",
				filepath: "cmd/diffence/main.go",
			},
			wantErr: wantErr{false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header, errHeader := extractHeader(tt.args.in)
			equals(t, tt.wantErr.header, errHeader != nil)
			equals(t, tt.want.header, header)
			filepath, errFilepath := extractFilePath(tt.args.in)
			equals(t, tt.wantErr.filepath, errFilepath != nil)
			equals(t, tt.want.filepath, filepath)
		})
	}
}
