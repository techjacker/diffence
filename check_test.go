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
	// rulesExtended := getRuleFile("test/fixtures/rules/rules_extended_regex.json")
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
			name: "Recognises an offensive diff - single_fail.diff",
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
			name: "Recognises an offensive diff - multi_fail.diff",
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
		// {
		// 	name: "Recognises an offensive diff - sqlpassword.diff",
		// 	args: args{
		// 		r:     getFixtureFile("test/fixtures/diffs/sqlpassword.diff"),
		// 		rules: rulesExtended,
		// 	},
		// 	want: Result{
		// 		Matched: true,
		// 		MatchedRules: MatchedRules{
		// 			"web/src/main/resources/db/migration/V0_2__SeedData.sql": []Rule{(*rulesExtended)[0]},
		// 			// "web/src/main/resources/db/migration/V0_4__AdminPassword.sql": *ruleSingle,
		// 		},
		// 	},
		// 	wantErr: false,
		// },
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
				t.Errorf("CheckDiffs(): %s error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckDiffs(): %s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
