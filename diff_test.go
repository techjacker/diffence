package diffence

import "testing"

type wantDiff struct {
	header   string
	filepath string
}

type wantErr struct {
	header   bool
	filepath bool
}

func TestDiffPush(t *testing.T) {
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
			name: "Diff.Push() - Admin Password",
			args: args{
				in: "diff --git a/web/src/main/resources/db/migration/V0_4__AdminPassword.sql b/web/src/main/resources/db/migration/V0_4__AdminPassword.sql" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want:    wantDiff{filepath: "web/src/main/resources/db/migration/V0_4__AdminPassword.sql"},
			wantErr: wantErr{false, false},
		},
		{
			name: "Diff.Push()",
			args: args{
				in: "diff --git a/lib/check.go b/lib/check.go" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want:    wantDiff{filepath: "lib/check.go"},
			wantErr: wantErr{false, false},
		},
		{
			name: "Diff.Push()",
			args: args{
				in: "diff --git a/README.md b/README.md" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want:    wantDiff{filepath: "README.md"},
			wantErr: wantErr{false, false},
		},
		{
			name: "Diff.Push()",
			args: args{
				in: "diff --git a/TODO.md b/TODO.md" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want:    wantDiff{filepath: "TODO.md"},
			wantErr: wantErr{false, false},
		},
		{
			name: "Diff.Push()",
			args: args{
				in: "hello world",
			},
			want:    wantDiff{filepath: ""},
			wantErr: wantErr{true, true},
		},
		{
			name: "Diff.Push()",
			args: args{
				in: "diff --git a/cmd/diffence/main.go b/cmd/diffence/main.go" +
					"\n" +
					"index 08044af..098342c 100644",
			},
			want:    wantDiff{filepath: "cmd/diffence/main.go"},
			wantErr: wantErr{false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filepath, errFilepath := extractFilePath(tt.args.in)
			equals(t, tt.wantErr.filepath, errFilepath != nil)
			equals(t, tt.want.filepath, filepath)
		})
	}
}
