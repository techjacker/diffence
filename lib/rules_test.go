package diffence

import (
	"reflect"
	"testing"
)

func Test_readRules(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    *[]rule
		wantErr bool
	}{
		{
			name: "Read rules from file",
			args: args{filePath: "./../test/fixtures/rules/rules.json"},
			want: &[]rule{
				{
					Caption:     "Contains word: password",
					Description: nil,
					Part:        "filename",
					Pattern:     "password",
					Type:        "regex",
				},
			},
			wantErr: false,
		},
		{
			name:    "Read rules from file",
			args:    args{filePath: "./../test/fixtures/does_not_exist.json"},
			want:    &[]rule{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadRulesFromFile(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadRulesFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadRulesFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
