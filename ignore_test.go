package diffence

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestIgnorer_Match(t *testing.T) {
	type fields struct {
		patterns []string
	}
	type args struct {
		in string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "Ignorer.Match() does find string in its list of patterns",
			fields: fields{getFileAsSlice("test/fixtures/ignore")},
			args:   args{"one"},
			want:   true,
		},
		{
			name:   "Ignorer.Match() does find string in its list of patterns",
			fields: fields{getFileAsSlice("test/fixtures/ignore")},
			args:   args{"not_in_pattern_list"},
			want:   false,
		},
		{
			name:   "Ignorer.Match() does find string in its list of patterns",
			fields: fields{[]string{"found*"}},
			args:   args{"found"},
			want:   true,
		},
		{
			name:   "Ignorer.Match() does find string in its list of patterns",
			fields: fields{[]string{"somestring"}},
			args:   args{"not_found"},
			want:   false,
		},
		{
			name:   "Ignorer.Match() does find string in its list of patterns",
			fields: fields{[]string{}},
			args:   args{"not_found"},
			want:   false,
		},
		{
			name:   "Ignorer.Match() does find string in its list of patterns",
			fields: fields{getFileAsSlice("/does/not/exist")},
			args:   args{"not_found"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Ignorer{
				patterns: tt.fields.patterns,
			}
			if got := i.Match(tt.args.in); got != tt.want {
				t.Errorf("Ignorer.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFileAsSlice(t *testing.T) {
	type args struct {
		fPath string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "getFileAsSlice() does find file",
			args: args{"test/fixtures/ignore"},
			want: []string{"one", "two", "three"},
		},
		{
			name: "getFileAsSlice() does not find file - so returns empty slice without panicing",
			args: args{"test/fixtures/does/not/exist"},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFileAsSlice(tt.args.fPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFileAsSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitLines(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "SplitDiffs() ",
			args: args{r: getFixtureFile("test/fixtures/ignore")},
			want: []string{
				"one",
				"two",
				"three",
			},
		},
		{
			name: "SplitDiffs() empty io.Reader",
			args: args{bytes.NewBuffer([]byte{})},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitLines(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
