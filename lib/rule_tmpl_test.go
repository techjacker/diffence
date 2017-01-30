package diffence

import "testing"

func Test_rule_Run(t *testing.T) {
	type fields struct {
		Caption     string
		Description interface{}
		Part        string
		Pattern     string
		Type        string
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
			name: "Should find a match",
			fields: fields{
				Caption:     "Contains word: password",
				Description: nil,
				Part:        "filename",
				Pattern:     "password",
				Type:        "regex",
			},
			args: args{"/path/to/password.txt"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rule{
				Caption:     tt.fields.Caption,
				Description: tt.fields.Description,
				Part:        tt.fields.Part,
				Pattern:     tt.fields.Pattern,
				Type:        tt.fields.Type,
			}
			if got := r.Run(tt.args.in); got != tt.want {
				t.Errorf("rule.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
