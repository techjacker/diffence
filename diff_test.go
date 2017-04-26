package diffence

import (
	"testing"
)

type wantErr struct {
	header bool
	fPath  bool
}

func TestDiffPush(t *testing.T) {
	type fields struct {
		ignorer Matcher
		Items   []DiffItem
		Error   error
	}
	type args struct {
		in string
	}
	tests := []struct {
		name   string
		fields fields
		want   DiffItem
		args   args
	}{
		{
			name:   "Diff.Push() - Commit ID Header",
			fields: fields{},
			args: args{
				in: "17949087b8e0c9179345e8dbb7b6705b49c93c77 Adds results logger" +
					"\n" +
					"diff --git a/web/src/main/resources/db/migration/V0_4__AdminPassword.sql b/web/src/main/resources/db/migration/V0_4__AdminPassword.sql" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want: DiffItem{
				commit: "17949087b8e0c9179345e8dbb7b6705b49c93c77 Adds results logger",
				fPath:  "web/src/main/resources/db/migration/V0_4__AdminPassword.sql",
			},
		},
		{
			name:   "Diff.Push() - Commit ID Header",
			fields: fields{},
			args: args{
				in: "diff --git a/web/src/main/resources/db/migration/V0_4__AdminPassword.sql b/web/src/main/resources/db/migration/V0_4__AdminPassword.sql" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want: DiffItem{
				commit: "",
				fPath:  "web/src/main/resources/db/migration/V0_4__AdminPassword.sql",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Diff{
				ignorer: tt.fields.ignorer,
				Items:   tt.fields.Items,
				Error:   tt.fields.Error,
			}
			d.Push(tt.args.in)
			equals(t, tt.want.fPath, d.Items[0].fPath)
			equals(t, tt.want.commit, d.Items[0].commit)
		})
	}
}

func TestExtractFilePath(t *testing.T) {
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
			name: "extractFilePath - Admin Password",
			args: args{
				in: "diff --git a/web/src/main/resources/db/migration/V0_4__AdminPassword.sql b/web/src/main/resources/db/migration/V0_4__AdminPassword.sql" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want:    DiffItem{fPath: "web/src/main/resources/db/migration/V0_4__AdminPassword.sql"},
			wantErr: wantErr{false, false},
		},
		{
			name: "extractFilePath",
			args: args{
				in: "diff --git a/lib/check.go b/lib/check.go" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want:    DiffItem{fPath: "lib/check.go"},
			wantErr: wantErr{false, false},
		},
		{
			name: "extractFilePath",
			args: args{
				in: "diff --git a/README.md b/README.md" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want:    DiffItem{fPath: "README.md"},
			wantErr: wantErr{false, false},
		},
		{
			name: "extractFilePath",
			args: args{
				in: "diff --git a/TODO.md b/TODO.md" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want:    DiffItem{fPath: "TODO.md"},
			wantErr: wantErr{false, false},
		},
		{
			name: "extractFilePath",
			args: args{
				in: "hello world",
			},
			want:    DiffItem{fPath: ""},
			wantErr: wantErr{true, true},
		},
		{
			name: "extractFilePath",
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

func Test_beginsWithHash(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "beginsWithHash - Commit ID Header",
			args: args{
				s: "17949087b8e0c9179345e8dbb7b6705b49c93c77 Adds results logger" +
					"\n" +
					"diff --git a/web/src/main/resources/db/migration/V0_4__AdminPassword.sql b/web/src/main/resources/db/migration/V0_4__AdminPassword.sql" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want: true,
		},
		{
			name: "beginsWithHash - Diff Header",
			args: args{
				s: "diff --git a/web/src/main/resources/db/migration/V0_4__AdminPassword.sql b/web/src/main/resources/db/migration/V0_4__AdminPassword.sql" +
					"\n" +
					"index 82366e3..5fc99b9 100644",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := beginsWithHash(tt.args.s); got != tt.want {
				t.Errorf("beginsWithHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
