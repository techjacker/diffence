package diffence

import (
	"bytes"
	"io"
	"log"
	"os"
	"path"
	"reflect"
	"testing"
)

func TestNewDiffer(t *testing.T) {
	tests := []struct {
		name string
		want Differ
	}{
		{
			name: "NewDiffer factory test",
			want: &diff{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDiffer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDiffer() = %v, want %v", got, tt.want)
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
	tests := []struct {
		name string
		args args
		want []DiffItem
	}{
		{
			name: "Differ.Parse()",
			args: args{
				r: getFixtureFile("test/fixtures/diffs/single.diff"),
			},
			want: []DiffItem{
				DiffItem{
					raw:      []byte{},
					filename: []byte("README.md"),
				},
			},
		},
		{
			name: "Differ.Parse()",
			args: args{
				r: getFixtureFile("test/fixtures/diffs/multi.diff"),
			},
			want: []DiffItem{
				DiffItem{
					raw:      []byte{},
					filename: []byte("TODO.md"),
				},
				DiffItem{
					raw:      []byte{},
					filename: []byte("systemdlogger/aws.py"),
				},
				DiffItem{
					raw:      []byte{},
					filename: []byte("systemdlogger/cloudwatch.py"),
				},
				DiffItem{
					raw:      []byte{},
					filename: []byte("tests/fixtures/config.json"),
				},
				DiffItem{
					raw:      []byte{},
					filename: []byte("tests/test_aws.py"),
				},
				DiffItem{
					raw:      []byte{},
					filename: []byte("tests/test_cloudwatch.py"),
				},
				DiffItem{
					raw:      []byte{},
					filename: []byte("tests/test_runner_integration.py"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			d := diff{}
			if err := d.Parse(tt.args.r); err != nil {
				t.Fatalf("diff.Parse() throw error %#v", err)

			}

			// check separating items correctly
			for _, di := range d.items {
				prefix := []byte("diff --git a")
				if !bytes.HasPrefix(di.raw, prefix) {
					t.Fatalf("diff.Parse() not separating items correctly \n\nGOT: %s, \n\nWANT to start with: %s", di.raw, prefix)
				}
			}
			if len(d.items) != len(tt.want) {
				t.Errorf("diff.Parse() \n\nGOT: %d items, \n\nWANT: %d items", len(d.items), len(tt.want))
			}

			// check lexing is correct
			if !reflect.DeepEqual(d.items[0].filename, tt.want[0].filename) {
				t.Errorf("diff.Parse() \n\nGOT: %s, \n\nWANT: %s", d.items[0].filename, tt.want[0].filename)
			}
		})
	}
}
