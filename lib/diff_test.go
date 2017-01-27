package diffence

import (
	"reflect"
	"testing"
)

func Test_extractFileName(t *testing.T) {
	type args struct {
		in string
	}
	type want struct {
		header   string
		filename string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "diff.getHeader()",
			args: args{
				in: "diff --git a/README.md b/README.md" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want: want{
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
			want: want{
				header:   "diff --git a/TODO.md b/TODO.md",
				filename: "TODO.md",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DiffItem{tt.args.in}
			if got := d.getHeader(); !reflect.DeepEqual(got, tt.want.header) {
				t.Errorf("d.getHeader()\nGOT:%s\nWANT:%s", got, tt.want.header)
			}
		})
	}
}
