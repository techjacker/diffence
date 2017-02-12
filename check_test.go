package diffence

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestCheckDiffs(t *testing.T) {
	type args struct {
		r     io.Reader
		rules *[]Rule
	}

	ruleSingle := getRuleFile("test/fixtures/rules/rules.json")
	ruleMulti := getRuleFile("test/fixtures/rules/rules_multi.json")

	tests := []struct {
		name    string
		args    args
		want    Result
		wantErr bool
	}{
		{
			name: "Recognises an offensive diff",
			args: args{
				r:     getFixtureFile("test/fixtures/diffs/single_fail.diff"),
				rules: ruleSingle,
			},
			want: Result{
				Matched: true,
				MatchedRules: MatchedRules{
					"path/to/password.txt": *ruleSingle,
				},
			},
			wantErr: false,
		},
		{
			name: "Recognises an offensive diff",
			args: args{
				r:     getFixtureFile("test/fixtures/diffs/single_fail.diff"),
				rules: ruleMulti,
			},
			want: Result{
				Matched: true,
				MatchedRules: MatchedRules{
					"path/to/password.txt": *ruleSingle,
				},
			},
			wantErr: false,
		},
		{
			name: "Recognises an offensive diff",
			args: args{
				r:     getFixtureFile("test/fixtures/diffs/multi_fail.diff"),
				rules: ruleMulti,
			},
			want: Result{
				Matched: true,
				MatchedRules: MatchedRules{
					"path/to/password.txt": []Rule{(*ruleMulti)[0]},
					"another/file/aws.pem": []Rule{(*ruleMulti)[1]},
				},
			},
			wantErr: false,
		},
		{
			name: "Recognises non diff text",
			args: args{
				r:     bytes.NewReader([]byte("not a diff")),
				rules: ruleMulti,
			},
			want: Result{
				Matched:      false,
				MatchedRules: MatchedRules{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dc := DiffChecker{tt.args.rules}
			got, err := dc.Check(tt.args.r)
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
