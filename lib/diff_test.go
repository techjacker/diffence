package diffence

import "testing"

type wantDiff struct {
	header   string
	filename string
}

func Test_extractFileName(t *testing.T) {
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
			d := DiffItem{tt.args.in}
			equals(t, d.getHeader(), tt.want.header)
			equals(t, d.getFilename(), tt.want.filename)
		})
	}
}
