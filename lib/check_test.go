package diffence

import (
	"io"
	"reflect"
	"testing"
)

func TestCheckDiffs(t *testing.T) {
	type args struct {
		r     io.Reader
		rules *[]rule
	}
	tests := []struct {
		name    string
		args    args
		want    Results
		wantErr bool
	}{
		{
			name: "Recognises an offensive diff",
			args: args{
				r: getFixtureFile("test/fixtures/diffs/single_fail.diff"),
				rules: &[]rule{
					rule{Caption: "Contains word: password",
						Description: nil,
						Part:        "filename",
						Pattern:     "password",
						Type:        "regex"},
				},
			},
			want: Results{
				"path/to/password.txt": []rule{
					rule{Caption: "Contains word: password",
						Description: nil,
						Part:        "filename",
						Pattern:     "password",
						Type:        "regex"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckDiffs(tt.args.r, tt.args.rules)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckDiffs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckDiffs() = %v, want %v", got, tt.want)
			}
		})
	}
}
