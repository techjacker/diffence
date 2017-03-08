package diffence

import "testing"

type wantErr struct {
	header bool
	fPath  bool
}

func TestDiffPush(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    DiffItem
		wantErr wantErr
	}{

		{
			name: "Diff.Push() - Admin Password",
			args: args{
				in: "diff --git a/web/src/main/resources/db/migration/V0_4__AdminPassword.sql b/web/src/main/resources/db/migration/V0_4__AdminPassword.sql" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want:    DiffItem{fPath: "web/src/main/resources/db/migration/V0_4__AdminPassword.sql"},
			wantErr: wantErr{false, false},
		},
		{
			name: "Diff.Push()",
			args: args{
				in: "diff --git a/lib/check.go b/lib/check.go" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want:    DiffItem{fPath: "lib/check.go"},
			wantErr: wantErr{false, false},
		},
		{
			name: "Diff.Push()",
			args: args{
				in: "diff --git a/README.md b/README.md" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want:    DiffItem{fPath: "README.md"},
			wantErr: wantErr{false, false},
		},
		{
			name: "Diff.Push()",
			args: args{
				in: "diff --git a/TODO.md b/TODO.md" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want:    DiffItem{fPath: "TODO.md"},
			wantErr: wantErr{false, false},
		},
		{
			name: "Diff.Push()",
			args: args{
				in: "hello world",
			},
			want:    DiffItem{fPath: ""},
			wantErr: wantErr{true, true},
		},
		{
			name: "Diff.Push()",
			args: args{
				in: "diff --git a/cmd/diffence/main.go b/cmd/diffence/main.go" +
					"\n" +
					"index 08044af..098342c 100644",
			},
			want:    DiffItem{fPath: "cmd/diffence/main.go"},
			wantErr: wantErr{false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fPath, errFilepath := extractFilePath(tt.args.in)
			equals(t, tt.wantErr.fPath, errFilepath != nil)
			equals(t, tt.want.fPath, fPath)
		})
	}
}
