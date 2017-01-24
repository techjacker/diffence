package diffence

import (
	"bytes"
	"io"
	"log"
	"os"
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

func Test_diff_Parse(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want []diffItem
	}{
		{
			name: "Differ.Parse()",
			// args: args{r: getFixtureFile("test/fixtures/diffs/single.diff")},
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
			want: []diffItem{
				diffItem{[]byte(
					"diff --git a/README.md b/README.md" +
						"\n" +
						"index 82366e3..5fc99b9 100644",
				)},
				diffItem{[]byte(
					"diff --git a/TODO.md b/TODO.md" +
						"\n" +
						"index 82366e3..5fc99b9 100644",
				)},
			},
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := diff{}
			if got := d.Parse(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diff.Parse() \n\nGOT: %s, \n\nWANT: %s", got, tt.want)
			}
		})
	}
}

func getFixtureFile(filename string) io.Reader {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
