package diffence

import (
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := diff{}
			d.Parse(tt.args.r)
			if !reflect.DeepEqual(d.items[0].filename, tt.want[0].filename) {
				t.Errorf("diff.Parse() \n\nGOT: %s, \n\nWANT: %s", d.items[0].filename, tt.want[0].filename)
			}
		})
	}
}
